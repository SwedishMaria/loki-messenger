<script lang="ts">
	// @ts-nocheck
	/**
	 * @component RegisterForm
	 * @description A Svelte component that provides a registration form for new users.
	 */
	import { goto } from '$app/navigation';
	let username = '';
	let password = '';

	/**
	 * @function register
	 * @description Handles the user registration process by sending a POST request to the backend.
	 * On success, it navigates to the login page.
	 */
	async function register() {
		const res = await fetch('/api/auth/register', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ username, password })
		});

		if (res.ok) {
			goto('/login');
		} else {
			alert('Failed to register');
		}
	}
</script>

<form on:submit|preventDefault={register}>
	<input type="text" bind:value={username} placeholder="Username" />
	<input type="password" bind:value={password} placeholder="Password" />
	<button type="submit">Register</button>
</form>

<style>
    form {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        width: 300px;
    }

    input {
        background-color: #000;
        color: #fff;
        border: 1px solid #fff;
        padding: 0.5rem;
        font-family: 'Press Start 2P', cursive;
    }

    button {
        background-color: #fff;
        color: #000;
        border: none;
        padding: 0.5rem;
        cursor: pointer;
        font-family: 'Press Start 2P', cursive;
    }
</style>