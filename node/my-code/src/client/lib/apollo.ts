import { ApolloClient, ApolloLink, HttpLink } from '@apollo/client';
import { onError } from '@apollo/client/link/error';
import { makeUniqueId } from '@apollo/client/utilities';
import { apolloCache } from './apollo-cache';
import { messageBroker } from './message-broker';

let graphqlUri = '';
if (process.env.NODE_ENV === 'production') {
    graphqlUri = 'https://otus-my-code.herokuapp.com/graphql';
} else {
    graphqlUri = 'http://localhost:3000/graphql';
}

const httpLink = new HttpLink({
    uri: graphqlUri
});

const errorLink = onError(({ graphQLErrors, networkError }) => {
    debugger;
    if (graphQLErrors)
        graphQLErrors.forEach(
            ({ message, locations, path }) => messageBroker.addErrorMessage(message)
            // console.log(
            //     `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
            // )
        );
    if (networkError) console.log(`[Network error]: ${networkError}`);
});

const apolloClient = new ApolloClient({
    cache: apolloCache,
    link: ApolloLink.from([errorLink, httpLink])
});

export default apolloClient;
