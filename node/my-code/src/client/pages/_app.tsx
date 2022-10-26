import { ApolloProvider } from '@apollo/client';

import apolloClient from '../lib/apollo';
import { Layout } from '../components/Layout';

import 'bootstrap/dist/css/bootstrap.min.css';

import '../styles/app.scss';

function MyApp({ Component, pageProps }) {
    return (
        <ApolloProvider client={apolloClient}>
            <Layout>
                <Component {...pageProps} />
            </Layout>
        </ApolloProvider>
    );
}

export default MyApp;
