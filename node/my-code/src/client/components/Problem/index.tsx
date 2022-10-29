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
        <div className={classnames(className, styles.problem)}>
            Problem Id: {problem._id}
            <br />
            Problem Summary: {problem.summary}
            <br />
            Problem Description: {problem.description}
            <br />
            Problem Solution: {problem.solution}
        </div>
    );
};

export { ProblemComponent as Problem };
