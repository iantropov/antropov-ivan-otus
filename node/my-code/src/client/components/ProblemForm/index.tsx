import React, { useState } from 'react';
import classnames from 'classnames';
import { useQuery } from '@apollo/client';

import { CategoriesReponse, Problem, ProblemData } from '../../lib/types';
import { GET_CATEGORIES_QUERY } from '../../lib/graphql';

import styles from './styles.module.scss';

interface ProblemFormProps {
    className?: string;
    problem?: Problem;
    onSubmit: (problem: ProblemData) => Promise<any>;
}

export const ProblemForm: React.FC<ProblemFormProps> = ({ className, problem, onSubmit }) => {
    const { data, loading } = useQuery<CategoriesReponse>(GET_CATEGORIES_QUERY);
    const categories = data?.categories;

    const [summary, setSummary] = useState(problem?.summary || '');
    const [description, setDescription] = useState(problem?.description || '');
    const [solution, setSolution] = useState(problem?.solution || '');
    const [categoryIds, setCategoryIds] = useState(problem?.categoryIds || []);
    const [isProblemCreating, setIsProblemCreating] = useState(false);

    if (loading) return <p>Loading...</p>;

    const onSummaryChange = event => {
        setSummary(event.currentTarget.value);
    };

    const onDescriptionChange = event => {
        setDescription(event.currentTarget.value);
    };

    const onSolutionChange = event => {
        setSolution(event.currentTarget.value);
    };

    const onCategoryIdsChange = event => {
        const newCategoryIds = [];
        [...event.currentTarget.options].forEach(option => {
            if (option.selected) {
                newCategoryIds.push(option.value);
            }
        });
        setCategoryIds(newCategoryIds);
    };

    const onFormSubmit = event => {
        event.preventDefault();

        setIsProblemCreating(true);

        onSubmit({
            summary,
            description,
            solution,
            categoryIds
        }).catch(error => {
            setIsProblemCreating(false);
            console.log(error);
            alert(error);
        });
    };

    return (
        <section className={classnames(className, styles.problemForm)}>
            <form
                className={classnames(styles.problemForm__rows, styles.problemFormRows)}
                onSubmit={onFormSubmit}
            >
                <div className={classnames(styles.problemFormRows__row, styles.problemFormRow)}>
                    <label className={classnames(styles.problemFormRow__label, 'form-label')}>
                        Summary:
                        <input
                            className="form-control"
                            type="text"
                            value={summary}
                            onChange={onSummaryChange}
                        />
                    </label>
                </div>
                <div className={classnames(styles.problemFormRows__row, styles.problemFormRow)}>
                    <label className={classnames(styles.problemFormRow__label, 'form-label')}>
                        Description:
                        <textarea
                            className="form-control"
                            onChange={onDescriptionChange}
                            value={description}
                        />
                    </label>
                </div>
                <div className={classnames(styles.problemFormRows__row, styles.problemFormRow)}>
                    <label className={classnames(styles.problemFormRow__label, 'form-label')}>
                        Solution:
                        <textarea
                            className="form-control"
                            onChange={onSolutionChange}
                            value={solution}
                        />
                    </label>
                </div>
                <div className={classnames(styles.problemFormRows__row, styles.problemFormRow)}>
                    <label className={classnames(styles.problemFormRow__label, 'form-label')}>
                        Categories:
                        <select
                            className="form-control"
                            onChange={onCategoryIdsChange}
                            multiple={true}
                            value={categoryIds}
                        >
                            {categories.map(category => (
                                <option key={category._id} value={category._id}>
                                    {category.name}
                                </option>
                            ))}
                        </select>
                    </label>
                </div>
                <div className={styles.problemFormRows__row}>
                    <button
                        type="submit"
                        className={classnames(styles.problemFormRow__button, 'btn', 'btn-primary')}
                        disabled={isProblemCreating || !summary || !description || !solution}
                    >
                        {isProblemCreating ? 'Saving...' : 'Save'}
                    </button>
                </div>
            </form>
        </section>
    );
};
