import React from 'react';
import classnames from 'classnames';

import { Problem } from '../../lib/types';

import styles from './styles.module.scss';

interface ProblemProps {
    className?: string;
    problem: Problem;
}

const ProblemComponent: React.FC<ProblemProps> = ({ className, problem }) => {
    return (
        <div
            className={classnames(
                className,
                styles.problem,
                'list-group-item list-group-item-action d-flex gap-3 py-3'
            )}
        >
            <img
                src="https://github.com/twbs.png"
                alt="twbs"
                width="32"
                height="32"
                className="rounded-circle flex-shrink-0"
            />
            <div className="d-flex gap-2 w-100 justify-content-between">
                <div>
                    <h6 className="mb-0">{problem.summary}</h6>
                    <p className={classnames(styles.problem__description)}>
                        {problem.description}
                    </p>
                    {problem.solution ? (
                        <>
                            <p>
                                <a
                                    className={styles.problem__link}
                                    data-bs-toggle="collapse"
                                    href={`#solution-${problem._id}`}
                                >
                                    Solution
                                </a>
                            </p>
                            <div className="collapse" id={`solution-${problem._id}`}>
                                <p className="mb-0 opacity-75">{problem.solution}</p>
                            </div>
                        </>
                    ) : null}
                </div>
                <small className="opacity-50 text-nowrap">#{problem._id}</small>
            </div>
        </div>
    );
};

export { ProblemComponent as Problem };
