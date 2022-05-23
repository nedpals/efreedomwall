import UUID from 'uuidjs';
import { ref } from 'vue';

const sessionIdKey = '__session_id';

export const sessionId = ref<string>(generateSessionId());

export function generateSessionId(): string {
    const gotSessionId = localStorage.getItem(sessionIdKey);
    if (!gotSessionId) {
        const newSessionId = UUID.generate();
        localStorage.setItem(sessionIdKey, newSessionId);
        return newSessionId;
    }
    return gotSessionId;
}