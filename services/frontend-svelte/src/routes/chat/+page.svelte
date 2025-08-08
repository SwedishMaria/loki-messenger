<script lang="ts">
	// @ts-nocheck
	/**
	 * @component ChatPage
	 * @description The main chat interface for sending and receiving encrypted messages.
	 */
	import { onMount } from 'svelte';
	import ScrambledText from '../../components/ScrambledText.svelte';
	import * as libsignal from '@privacyresearch/libsignal-protocol-typescript';
	import { initSignal } from '../../lib/signal';

	let socket: WebSocket;
	let messages: string[] = [];
	let message = '';
	let aliceCipher: libsignal.SessionCipher;
	let bobCipher: libsignal.SessionCipher;

	/**
	 * @function arrayBufferToString
	 * @description Converts an ArrayBuffer to a string.
	 * @param buffer The ArrayBuffer to convert.
	 * @returns The converted string.
	 */
	function arrayBufferToString(buffer: ArrayBuffer) {
		return String.fromCharCode.apply(null, new Uint8Array(buffer) as any);
	}

	/**
	 * @function stringToArrayBuffer
	 * @description Converts a string to an ArrayBuffer.
	 * @param str The string to convert.
	 * @returns The converted ArrayBuffer.
	 */
	function stringToArrayBuffer(str: string) {
		const buf = new ArrayBuffer(str.length);
		const bufView = new Uint8Array(buf);
		for (let i = 0, strLen = str.length; i < strLen; i++) {
			bufView[i] = str.charCodeAt(i);
		}
		return buf;
	}

	onMount(async () => {
		// Initialize the Signal protocol stores and ciphers for Alice and Bob.
		const { aliceStore, bobStore } = await initSignal();
		aliceCipher = new libsignal.SessionCipher(aliceStore, new libsignal.SignalProtocolAddress('bob', 1));
		bobCipher = new libsignal.SessionCipher(bobStore, new libsignal.SignalProtocolAddress('alice', 1));

		// Establish a WebSocket connection.
		const token = localStorage.getItem('token');
		socket = new WebSocket(`ws://localhost:8080/api/ws?token=${token}`);

		// Handle incoming messages from the WebSocket.
		socket.onmessage = async (event) => {
			const encrypted = JSON.parse(event.data);
			let decrypted: ArrayBuffer;

			// Decrypt the message using the appropriate Signal protocol method.
			if (encrypted.type === 3) {
				decrypted = await bobCipher.decrypt(libsignal.PreKeyWhisperMessage.deserialize(stringToArrayBuffer(encrypted.body)));
			} else {
				decrypted = await bobCipher.decrypt(libsignal.WhisperMessage.deserialize(stringToArrayBuffer(encrypted.body)));
			}
			messages = [...messages, new TextDecoder().decode(decrypted)];
		};
	});

	/**
	 * @function sendMessage
	 * @description Encrypts and sends a message through the WebSocket.
	 */
	async function sendMessage() {
		if (socket && message) {
			// Encrypt the message using the Signal protocol.
			const encrypted = await aliceCipher.encrypt(new TextEncoder().encode(message).buffer);
			// Send the encrypted message to the server.
			socket.send(JSON.stringify({ type: encrypted.type, body: arrayBufferToString(encrypted.serialize()) }));
			message = '';
		}
	}
</script>

<h1>Chat</h1>

<div class="chat-container">
    <div class="messages">
        {#each messages as msg, i (i)}
            <div class="message">
                <ScrambledText text={msg} />
            </div>
        {/each}
    </div>
    <form on:submit|preventDefault={sendMessage}>
        <input type="text" bind:value={message} placeholder="Type a message..." />
        <button type="submit">Send</button>
    </form>
</div>

<style>
    .chat-container {
        display: flex;
        flex-direction: column;
        height: 80vh;
        width: 100%;
        max-width: 800px;
        margin: 0 auto;
        border: 1px solid #fff;
    }

    .messages {
        flex-grow: 1;
        padding: 1rem;
        overflow-y: auto;
    }

    form {
        display: flex;
        padding: 1rem;
        gap: 1rem;
    }

    input {
        flex-grow: 1;
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