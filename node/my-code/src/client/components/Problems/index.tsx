import React from 'react';
import classnames from 'classnames';

import { Problem } from '../../lib/types';
import { Problem as ProblemComponent } from '../Problem';

import styles from './styles.module.scss';

interface ProblemsProps {
    className?: string;
    problems: Problem[];
    favorites: string[];
    allowEdit: boolean;
    allowDelete: boolean;
    hasNextPage: boolean;
    onFetchMore: () => void;
    onDelete?: (problem: Problem) => Promise<void>;
    onLike?: (problem: Problem) => Promise<void>;
    onUnlike?: (problem: Problem) => Promise<void>;
}

export const Problems: React.FC<ProblemsProps> = ({
    className,
    problems,
    favorites,
    allowEdit,
    allowDelete,
    hasNextPage,
    onFetchMore,
    onDelete,
    onLike,
    onUnlike
}) => {
    if (problems.length === 0) {
        return (
            <div className={classnames(className, styles.problems, styles.problems_empty)}>
                You don&apos;t have any problems yet :(
            </div>
        );
    }

    return (
        <div className={classnames(className, styles.problems)}>
            <ul className={classnames(className, styles.problems__list, 'list-group')}>
                {problems.map(problem => (
                    <li className={styles.problems__problem} key={problem._id}>
                        <ProblemComponent
                            problem={problem}
                            isLiked={favorites.includes(problem._id)}
                            allowEdit={allowEdit}
                            allowDelete={allowDelete}
                            onDelete={onDelete}
                            onLike={onLike}
                            onUnlike={onUnlike}
                        />
                    </li>
                ))}
            </ul>
            {hasNextPage && (
                <div className={classnames(styles.problems__more)}>
                    <button className="btn btn-primary" onClick={onFetchMore}>Fetch More</button>
                </div>
            )}
        </div>
    );
};
