import React, { useMemo } from 'react';
import classnames from 'classnames';
import Link from 'next/link';
import { useMutation } from '@apollo/client';

import { Problem } from '../../lib/types';
import {
    DELETE_PROBLEM_MUTATION,
    GET_PROBLEMS_QUERY,
    LIKE_PROBLEM_MUTATION,
    UNLIKE_PROBLEM_MUTATION,
    WHO_AM_I_QUERY
} from '../../lib/graphql';
import { useUser } from '../../hooks/use-user';

import styles from './styles.module.scss';

interface ProblemProps {
    className?: string;
    problem: Problem;
    allowEdit: boolean;
    allowRemove: boolean;
}

const ProblemComponent: React.FC<ProblemProps> = ({
    className,
    problem,
    allowEdit,
    allowRemove
}) => {
    const [user] = useUser();

    const [deleteProblem] = useMutation(DELETE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: GET_PROBLEMS_QUERY }]
    });

    const [likeProblem, { loading: isProblemLiking }] = useMutation(LIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });
    const [unlikeProblem, { loading: isProblemUnliking }] = useMutation(UNLIKE_PROBLEM_MUTATION, {
        refetchQueries: [{ query: WHO_AM_I_QUERY }]
    });

    const onDeleteProblemClick = () => {
        deleteProblem({ variables: { problemId: problem._id } }).then(
            () => {
                console.log(`Removed problem #${problem._id} successfully!`);
            },
            error => {
                alert(error);
            }
        );
    };

    const onLikeProblemClick = () => {
        (isLiked ? unlikeProblem : likeProblem)({ variables: { problemId: problem._id } }).then(
            () => {
                console.log(`${isLiked ? 'Unliked' : 'Liked'} problem #${problem._id} successfully!`);
            },
            error => {
                alert(error);
            }
        );
    };

    const isLiked = useMemo(() => {
        return user.favorites.includes(problem._id);
    }, [user.favorites.length]);

    return (
        <div
            className={classnames(
                className,
                styles.problem,
                'list-group-item list-group-item-action d-flex gap-3 py-3'
            )}
        >
            <div className={styles.problem__content}>
                <div className={styles.problem__header}>
                    <div className={styles.problem__icon}>
                        <img
                            src="https://github.com/twbs.png"
                            alt="twbs"
                            width="32"
                            height="32"
                            className="rounded-circle flex-shrink-0"
                        />
                    </div>
                    <div className={styles.problem__summary}>{problem.summary}</div>
                    <div className={styles.problem__id}>
                        <small className="opacity-50 text-nowrap">#{problem._id}</small>
                    </div>
                </div>
                <p className={styles.problem__categories}>
                    {problem.categories.map(({ name }) => `@${name}`).join(', ')}
                </p>
                <p className={styles.problem__description}>{problem.description}</p>
                <div className={styles.problem__solution}>
                    <p>
                        <a
                            className={styles.problem__collapse}
                            data-bs-toggle="collapse"
                            href={`#solution-${problem._id}`}
                        >
                            Solution
                        </a>
                    </p>
                    <div className="collapse" id={`solution-${problem._id}`}>
                        <p className="mb-0 opacity-75">{problem.solution}</p>
                    </div>{' '}
                </div>
            </div>
            <div className={styles.problem__footer}>
                <button
                    className={classnames(
                        styles.problem__button,
                        `btn btn-sm ${isLiked ? 'btn-warning' : 'btn-success'}`
                    )}
                    disabled={isProblemLiking || isProblemUnliking}
                    onClick={onLikeProblemClick}
                >
                    {isLiked ? 'Unlike' : 'Like'}
                </button>
                {allowEdit && (
                    <Link href={`/problems/${problem._id}`}>
                        <a className={classnames(styles.problem__button, 'btn btn-sm btn-primary')}>
                            Edit
                        </a>
                    </Link>
                )}
                {allowRemove && (
                    <button
                        className={classnames(styles.problem__button, 'btn btn-sm btn-danger')}
                        onClick={onDeleteProblemClick}
                    >
                        Delete
                    </button>
                )}
            </div>
        </div>
    );
};

export { ProblemComponent as Problem };
