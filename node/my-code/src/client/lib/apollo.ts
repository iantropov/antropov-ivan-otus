import { ApolloClient, ApolloLink, HttpLink } from '@apollo/client';
import { onError } from '@apollo/client/link/error';

import { apolloCache } from './apollo-cache';
import { messageBroker } from './message-broker';
import { GraphQLExtensions } from './types';

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
    if (graphQLErrors)
        graphQLErrors.forEach(({ message, extensions }) => {
            let errorMessage = message;
            const responseMessage = (extensions as GraphQLExtensions).response?.message;
            if (Array.isArray(responseMessage) && responseMessage.length > 0) {
                errorMessage = `${message}: ${responseMessage.join(', ')}`;
            } else if (responseMessage && responseMessage !== errorMessage) {
                errorMessage = `${message}: ${responseMessage}`
            }
            messageBroker.addErrorMessage(errorMessage);
        });
    if (networkError) messageBroker.addErrorMessage(`[Network error]: ${networkError}`);
});

const apolloClient = new ApolloClient({
    cache: apolloCache,
    link: ApolloLink.from([errorLink, httpLink])
});

export default apolloClient;
