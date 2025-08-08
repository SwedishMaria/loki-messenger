<script lang="ts">
	// @ts-nocheck
	/**
	 * @component LoginForm
	 * @description A Svelte component that provides a login form for user authentication.
	 */
	import { goto } from '$app/navigation';
	let username = '';
	let password = '';

	/**
	 * @function login
	 * @description Handles the login process by sending a POST request to the backend.
	 * On success, it stores the JWT in localStorage and navigates to the chat page.
	 */
	async function login() {
		const res = await fetch('/api/auth/login', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ username, password })
		});

		if (res.ok) {
			const data = await res.json();
			localStorage.setItem('token', data.token);
			goto('/chat');
		} else {
			alert('Invalid credentials');
		}
	}
</script>

<form on:submit|preventDefault={login}>
	<input type="text" bind:value={username} placeholder="Username" />
	<input type="password" bind:value={password} placeholder="Password" />
	<button type="submit">Login</button>
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