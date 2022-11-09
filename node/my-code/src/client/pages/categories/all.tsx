import React from 'react';
import { NextPage } from 'next';
import { useQuery } from '@apollo/client';
import classnames from 'classnames';
import Link from 'next/link';

import { useUser } from '../../hooks/use-user';
import { CategoriesResponse } from '../../lib/types';
import { GET_CATEGORIES_QUERY } from '../../lib/graphql';

import styles from './all.module.scss';

const Users: NextPage = () => {
    const { data, loading } = useQuery<CategoriesResponse>(GET_CATEGORIES_QUERY);

    const [user, isUserLoading] = useUser({ isAdmin: true });

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    return (
        <section className={classnames(styles.categories)}>
            <h2 className={classnames(styles.categories__header)}>Available categories:</h2>
            <table className={classnames(styles.categories__table, 'table')}>
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">Name</th>
                    </tr>
                </thead>
                <tbody>
                    {data.categories.map(category => (
                        <tr key={category._id}>
                            <th scope="row">{category._id}</th>
                            <td>{category.name}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
            <div className={classnames(styles.categories__buttons)}>
                <Link href="/categories/new">
                    <a className={classnames(styles.categories__createCategory, 'btn btn-primary')}>
                        Create Category
                    </a>
                </Link>
            </div>
        </section>
    );
};

export default Users;
