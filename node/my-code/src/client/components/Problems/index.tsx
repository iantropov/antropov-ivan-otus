import React from 'react';
import classnames from 'classnames';

import { Problem } from '../../lib/types';
import { Problem as ProblemComponent } from '../Problem';

import styles from './styles.module.scss';

interface ProblemsProps {
    className?: string;
    problems: Problem[];
    allowEdit: boolean;
    allowRemove: boolean;
}

export const Problems: React.FC<ProblemsProps> = ({
    className,
    problems,
    allowEdit,
    allowRemove
}) => {
    if (problems.length === 0) {
        return (
            <div className={classnames(className, styles.problems, styles.problems_empty)}>
                You don't have any problems yet :(
            </div>
        );
    }

    return (
        <ul className={classnames(className, styles.problems, 'list-group')}>
            {problems.map(problem => (
                <li className={styles.problems__problem} key={problem._id}>
                    <ProblemComponent
                        problem={problem}
                        allowEdit={allowEdit}
                        allowRemove={allowRemove}
                    />
                </li>
            ))}
        </ul>
    );
};
