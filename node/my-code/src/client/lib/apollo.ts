import { ApolloClient, InMemoryCache } from '@apollo/client';

let graphqlUri = '';
if (process.env.NODE_ENV === 'production') {
    graphqlUri = 'https://otus-my-code.herokuapp.com/graphql';
} else {
    graphqlUri = 'http://localhost:3000/graphql';
}

const apolloClient = new ApolloClient({
    uri: graphqlUri,
    cache: new InMemoryCache({
        typePolicies: {
            Query: {
                fields: {
                    searchProblems: {
                        keyArgs: ['text', 'categoryIds', 'favorites'],

                        merge(existing, incoming, { args: { cursor }, readField }) {
                            if (!cursor) {
                                return incoming;
                            }

                            const merged = existing ? existing.edges.slice(0) : [];

                            let offset = offsetFromCursor(merged, cursor, readField);
                            if (offset < 0) offset = merged.length;

                            for (let i = 0; i < incoming.edges.length; ++i) {
                                merged[offset + i] = incoming.edges[i];
                            }
                            return {
                                edges: merged,
                                pageInfo: incoming.pageInfo
                            };
                        }
                    }
                }
            }
        }
    })
});

function offsetFromCursor(items, cursor, readField) {
    for (let i = items.length - 1; i >= 0; --i) {
        const item = items[i];
        if (readField('_id', item) === cursor) {
            return i + 1;
        }
    }
    return -1;
}

export default apolloClient;
