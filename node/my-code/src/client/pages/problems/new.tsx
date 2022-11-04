import React from 'react';
import { NextPage } from 'next';
import { useMutation, useQuery } from '@apollo/client';
import { useRouter } from 'next/router';

import { CREATE_PROBLEM_MUTATION, GET_CATEGORIES_QUERY, GET_PROBLEMS_QUERY } from '../../lib/graphql';
import { CategoriesReponse, ProblemData } from '../../lib/types';
import { ProblemForm } from '../../components/ProblemForm';

import styles from './new.module.scss';

const NewProblem: NextPage = () => {
    const router = useRouter();

    const [createProblem] = useMutation(CREATE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: GET_PROBLEMS_QUERY }]
    });

    const onSubmit = (problem: ProblemData) => {
        return createProblem({ variables: problem }).then(() => {
            console.log(`Created a problem successfully!`);
            router.push('/problems/all');
        });
    };

    return (
        <main className={styles.newProblem}>
            <div className={styles.newProblem__header}>
                <h1>Create New Problem</h1>
            </div>
            <ProblemForm onSubmit={onSubmit} />
        </main>
    );
};

export default NewProblem;
