import { redirect } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async () => {
    if (typeof window !== 'undefined') {
        const token = localStorage.getItem('token');
        if (!token) {
            throw redirect(307, '/login');
        }
    }
    return {};
};