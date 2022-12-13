import { useQuery } from '@apollo/client';
import { useRouter } from 'next/router';
import { useEffect } from 'react';

import { WHO_AM_I_QUERY } from '../lib/graphql';
import { User, WhoAmIResponse } from '../lib/types';

interface UseUserOptions {
    isAdmin?: boolean;
}

export function useUser(options?: UseUserOptions): [User | null, boolean] {
    const router = useRouter();
    const { data: userData, loading: userLoading } = useQuery<WhoAmIResponse>(WHO_AM_I_QUERY);
    options = options ?? { isAdmin: false };

    useEffect(() => {
        if (!userLoading && (!userData?.whoAmI || (options.isAdmin && !userData?.whoAmI.isAdmin))) {
            router.replace('/users/login');
        }
    }, [userData, userLoading]);

    return [userData ? userData.whoAmI : null, userLoading];
}
