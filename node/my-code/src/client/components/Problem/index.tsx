import React, { useState } from 'react';
import classnames from 'classnames';
import Link from 'next/link';
import Image from 'next/image';

import { Problem } from '../../lib/types';

import styles from './styles.module.scss';

interface ProblemProps {
    className?: string;
    problem: Problem;
    isLiked: boolean;
    allowEdit: boolean;
    allowDelete: boolean;
    onDelete?: (problem: Problem) => Promise<void>;
    onLike?: (problem: Problem) => Promise<void>;
    onUnlike?: (problem: Problem) => Promise<void>;
}

const ProblemComponent: React.FC<ProblemProps> = ({
    className,
    problem,
    isLiked,
    allowEdit,
    allowDelete,
    onDelete,
    onLike,
    onUnlike
}) => {
    const [isProcessing, setIsProcessing] = useState(false);

    const processProblem = (processCallback?: (problem: Problem) => Promise<void>) => {
        return () => {
            if (!processCallback) return;

            setIsProcessing(true);
            processCallback(problem).finally(() => setIsProcessing(false));
        };
    };

    const onDeleteProblem = processProblem(onDelete);
    const onLikeProblem = processProblem(onLike);
    const onUnlikeProblem = processProblem(onUnlike);

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
                        <Image
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
                {problem.solution && (
                    <div className={styles.problem__solution}>
                        <a
                            className={styles.problem__collapse}
                            data-bs-toggle="collapse"
                            href={`#solution-${problem._id}`}
                        >
                            Solution
                        </a>
                        <div className="collapse" id={`solution-${problem._id}`}>
                            <p className="mb-0 opacity-75">{problem.solution}</p>
                        </div>
                    </div>
                )}
            </div>
            <div className={styles.problem__footer}>
                <button
                    className={classnames(
                        styles.problem__button,
                        `btn btn-sm ${isLiked ? 'btn-warning' : 'btn-success'}`
                    )}
                    disabled={isProcessing}
                    onClick={isLiked ? onUnlikeProblem : onLikeProblem}
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
                {allowDelete && (
                    <button
                        className={classnames(styles.problem__button, 'btn btn-sm btn-danger')}
                        onClick={onDeleteProblem}
                        disabled={isProcessing}
                    >
                        Delete
                    </button>
                )}
            </div>
        </div>
    );
};

export { ProblemComponent as Problem };
