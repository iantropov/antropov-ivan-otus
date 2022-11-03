import React from 'react';
import { NextPage } from 'next';
import { useMutation } from '@apollo/client';
import { useRouter } from 'next/router';

import { CREATE_PROBLEM_MUTATION, GET_PROBLEMS_QUERY } from '../../lib/graphql';
import { ProblemData } from '../../lib/types';
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
        <section className={styles.register}>
            <div className={styles.register__header}>
                <h1>Create New Problem</h1>
            </div>
            <ProblemForm onSubmit={onSubmit} />
        </section>
    );
};

export default NewProblem;
