<script lang="ts">
    import { onMount } from 'svelte';

    export let text: string;
    let displayedText = '';

    const chars = '!<>-_\\/[]{}—=+*^?#________';

    onMount(() => {
        let i = 0;
        const interval = setInterval(() => {
            displayedText = text
                .split('')
                .map((char, index) => {
                    if (index < i) {
                        return text[index];
                    }
                    return chars[Math.floor(Math.random() * chars.length)];
                })
                .join('');

            if (i >= text.length) {
                clearInterval(interval);
            }

            i += 1 / 3;
        }, 30);
    });
</script>

<span>{displayedText}</span>