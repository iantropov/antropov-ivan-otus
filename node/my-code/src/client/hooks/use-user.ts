import { useQuery } from '@apollo/client';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { WHO_AM_I_QUERY } from '../lib/graphql';
import { User, WhoAmIResponse } from '../lib/types';

export function useUser(): [User | null, boolean] {
    const router = useRouter();
    const { data: userData, loading: userLoading } = useQuery<WhoAmIResponse>(WHO_AM_I_QUERY);

    useEffect(() => {
        if (!userLoading && !userData?.whoAmI) {
            router.replace('/login');
        }
    }, [userData, userLoading]);

    return [userData ? userData.whoAmI : null, userLoading];
}
