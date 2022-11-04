import React, { useState } from 'react';
import classnames from 'classnames';
import { useQuery } from '@apollo/client';

import { CategoriesReponse, Problem, ProblemsSearchFilter } from '../../lib/types';
import { GET_CATEGORIES_QUERY } from '../../lib/graphql';

import styles from './styles.module.scss';

interface ProblemsSearchFilterProps {
    className?: string;
    onApply: (filter: ProblemsSearchFilter) => void;
}

const ProblemsSearchFilterComponent: React.FC<ProblemsSearchFilterProps> = ({
    className,
    onApply
}) => {
    const { data } = useQuery<CategoriesReponse>(GET_CATEGORIES_QUERY);
    const categories = data?.categories;

    const [text, setText] = useState('');
    const [categoryIds, setCategoryIds] = useState<string[]>([]);
    const [favorites, setFavorites] = useState(false);

    const onTextChange = event => {
        setText(event.currentTarget.value);
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

    const onFavoritesChange = event => {
        setFavorites(event.currentTarget.checked);
    };

    const onSubmit = event => {
        event.preventDefault();
        onApply({
            text,
            categoryIds,
            favorites
        });
    };

    return (
        <section className={classnames(className, styles.problemsSearchFilter, 'card')}>
            <form
                className={classnames(className, styles.problemsSearchFilter__form, 'card-body')}
                onSubmit={onSubmit}
            >
                <h2 className={styles.problemsSearchFilter__header}>Search Filter:</h2>
                <div className={classnames(styles.problemsSearchFilter__top)}>
                    <label className={classnames(styles.problemsSearchFilter__text, 'form-label')}>
                        Text:
                        <input
                            className={classnames('form-control')}
                            onChange={onTextChange}
                            type="search"
                            value={text}
                        />
                    </label>
                    <button className="btn btn-primary" type="submit">
                        Apply
                    </button>
                </div>
                <div className={classnames(styles.problemsSearchFilter__bottom)}>
                    <a
                        className={styles.problemsSearchFilter__more}
                        data-bs-toggle="collapse"
                        href="#more-filters"
                    >
                        More filters:
                    </a>
                    <div
                        className={classnames(styles.problemsSearchFilter__moreBody, 'collapse')}
                        id="more-filters"
                    >
                        <div
                            className={classnames(
                                styles.problemsSearchFilter__favorites,
                                'form-check'
                            )}
                        >
                            <input
                                className="form-check-input"
                                onChange={onFavoritesChange}
                                checked={favorites}
                                type="checkbox"
                                value="checked"
                                id="only-favorites"
                            />
                            <label className="form-check-label" htmlFor="only-favorites">
                                Default checkbox
                            </label>
                        </div>
                        <label
                            className={classnames(styles.problemsSearchFilter__label, 'form-label')}
                        >
                            Categories:
                            <select
                                className="form-control"
                                onChange={onCategoryIdsChange}
                                multiple={true}
                                value={categoryIds}
                            >
                                {categories?.map(category => (
                                    <option key={category._id} value={category._id}>
                                        {category.name}
                                    </option>
                                ))}
                            </select>
                        </label>
                    </div>
                </div>
            </form>
        </section>
    );
};

export { ProblemsSearchFilterComponent as ProblemsSearchFilter };
