import React, { useState } from 'react';
import { NextPage } from 'next';
import { useMutation } from '@apollo/client';
import { useRouter } from 'next/router';
import classnames from 'classnames';

import { CREATE_CATEGORY_MUTATION, GET_CATEGORIES_QUERY } from '../../lib/graphql';
import { Main } from '../../components/Main';
import { messageBroker } from '../../lib/message-broker';

import styles from './new.module.scss';

const NewCategory: NextPage = () => {
    const router = useRouter();

    const [name, setName] = useState('');
    const [isCreating, setIsCreating] = useState(false);

    const [createCategory] = useMutation(CREATE_CATEGORY_MUTATION, {
        refetchQueries: [{ query: GET_CATEGORIES_QUERY }]
    });

    const onNameChange = event => {
        setName(event.currentTarget.value);
    };

    const onSubmit = event => {
        event.preventDefault();
        setIsCreating(true);

        createCategory({ variables: { name } }).then(
            () => {
                router.push('/categories/all');
                messageBroker.addSuccessMessage('Created category successfully!');
            },
            () => {
                setIsCreating(false);
            }
        );
    };

    return (
        <Main className={styles.newCategory}>
            <div className={styles.newCategory__header}>
                <h1>New Category</h1>
            </div>
            <div className={classnames(styles.newCategory__body, styles.newCategoryBody)}>
                <form onSubmit={onSubmit}>
                    <div className={classnames(styles.newCategoryBody__row, styles.newCategoryRow)}>
                        <label className={classnames(styles.newCategoryRow__label, 'form-label')}>
                            Name:
                            <input
                                className="form-control"
                                type="text"
                                value={name}
                                onChange={onNameChange}
                            />
                        </label>
                    </div>
                    <div className={styles.newCategoryBody__row}>
                        <button
                            type="submit"
                            className={classnames(
                                styles.newCategoryRow__button,
                                'btn',
                                'btn-primary'
                            )}
                            disabled={isCreating || !name}
                        >
                            {isCreating ? 'Creating...' : 'Create'}
                        </button>
                    </div>
                </form>
            </div>
        </Main>
    );
};

export default NewCategory;
