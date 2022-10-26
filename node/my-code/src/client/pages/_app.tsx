import { ApolloProvider } from '@apollo/client';

import apolloClient from '../lib/apollo';

import "bootstrap/dist/css/bootstrap.min.css";

import '../styles/app.scss';
import '../styles/login.scss';

function MyApp({ Component, pageProps }) {
    return (
        <ApolloProvider client={apolloClient}>
            <Component {...pageProps} />
        </ApolloProvider>
    );
}

export default MyApp;
