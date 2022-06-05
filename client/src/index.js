import React from 'react';
import { render } from 'react-dom';
import App from './App';
import { ApolloClient, from, InMemoryCache, createHttpLink } from "@apollo/client/core";

const httpLink = createHttpLink({
    // https://oss.navercorp.com/ClouS/tower/pull/15#issue-2119194
    uri: '/api/graphql/query',
});

export const client = new ApolloClient({
    // https://www.apollographql.com/docs/react/networking/advanced-http-networking/
    link: from([
        httpLink,
    ]),
    cache: new InMemoryCache(),
    // https://www.apollographql.com/docs/react/api/core/ApolloClient#example-defaultoptions-object
    defaultOptions: {
        query: {
            fetchPolicy: 'network-only',
        },
    },
});

render(<App />, document.getElementById('app'));
