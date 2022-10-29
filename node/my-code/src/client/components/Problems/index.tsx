import React from 'react';
import classnames from 'classnames';

import { Problem } from '../../lib/types';
import { Problem as ProblemComponent } from '../Problem';

import styles from './styles.module.scss';

interface ProblemsProps {
    className?: string;
    problems: Problem[];
}

export const Problems: React.FC<ProblemsProps> = ({ className, problems }) => {
    return (
        <ul className={classnames(className, styles.problems)}>
            {problems.map(problem => (
                <li className={styles.problems__problem} key={problem._id}>
                    <ProblemComponent problem={problem} />
                </li>
            ))}
        </ul>
    );
};
