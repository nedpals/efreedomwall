import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';

dayjs.extend(relativeTime);

export function humanizeTime(dt: string | Date): string {
    return dayjs(dt).fromNow();
}