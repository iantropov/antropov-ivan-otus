import React from 'react';
import { NextPage } from 'next';
import { useMutation, useQuery } from '@apollo/client';
import { useRouter } from 'next/router';

import { UPDATE_PROBLEM_MUTATION, GET_PROBLEMS_QUERY, GET_PROBLEM_QUERY } from '../../lib/graphql';
import { ProblemData, ProblemReponse } from '../../lib/types';
import { ProblemForm } from '../../components/ProblemForm';
import { useUser } from '../../hooks/use-user';

interface EditProblemProps {
    id: string;
}

const EditProblemPage: NextPage = () => {
    const router = useRouter();
    const { id } = router.query;

    if (!id) {
        return null;
    }

    return <EditProblem id={id as string} />;
};

const EditProblem: React.FC<EditProblemProps> = ({ id }) => {
    const router = useRouter();

    const { data, loading } = useQuery<ProblemReponse>(GET_PROBLEM_QUERY, {
        variables: {
            id
        }
    });
    const [updateProblem] = useMutation(UPDATE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: GET_PROBLEMS_QUERY }]
    });

    const [user, isUserLoading] = useUser({ isAdmin: true });

    if (loading || isUserLoading) return <p>Loading...</p>;
    if (!user) return null;

    const onSubmit = (problem: ProblemData) => {
        return updateProblem({ variables: { id, ...problem } }).then(() => {
            console.log(`Updated problem #${id} successfully!`);
            router.push('/problems/all');
        });
    };

    return (
        <section>
            <div>
                <h1>Edit Problem: {data.problem.summary}</h1>
            </div>
            <ProblemForm problem={data.problem} onSubmit={onSubmit} />
        </section>
    );
};

export default EditProblemPage;
