const baseUrl = import.meta.env.API_URL || 'http://localhost:4003/api';

export function client(path: string, options?: RequestInit) {
    const endpoint = path.length > 0 && path[0] != '/' ? '/' + path : path;
    return fetch(`${baseUrl}${endpoint}`, options);
}