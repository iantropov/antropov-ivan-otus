import React, { useEffect } from 'react';
import { ApolloProvider } from '@apollo/client';

import apolloClient from '../lib/apollo';
import { Layout } from '../components/Layout';

import 'bootstrap/dist/css/bootstrap.min.css';

import '../styles/app.scss';

function MyApp({ Component, pageProps }) {
    useEffect(() => {
        require('bootstrap/dist/js/bootstrap.bundle.min.js');
    }, []);

    return (
        <ApolloProvider client={apolloClient}>
            <Layout>
                <Component {...pageProps} />
            </Layout>
        </ApolloProvider>
    );
}

export default MyApp;
