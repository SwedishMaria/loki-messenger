import * as libsignal from '@privacyresearch/libsignal-protocol-typescript';

const ALICE_ADDRESS = new libsignal.SignalProtocolAddress('alice', 1);
const BOB_ADDRESS = new libsignal.SignalProtocolAddress('bob', 1);

async function createDummySignalProtocolStore() {
    const store = new libsignal.SessionInMemoryStore();
    const identityKeyPair = await libsignal.KeyHelper.generateIdentityKeyPair();
    const registrationId = await libsignal.KeyHelper.generateRegistrationId();
    store.put('identityKey', identityKeyPair);
    store.put('registrationId', registrationId);
    return store;
}

export async function initSignal() {
    const aliceStore = await createDummySignalProtocolStore();
    const bobStore = await createDummySignalProtocolStore();

    const aliceSessionBuilder = new libsignal.SessionBuilder(aliceStore, BOB_ADDRESS);
    const bobSessionBuilder = new libsignal.SessionBuilder(bobStore, ALICE_ADDRESS);

    const bobIdentity = (await bobStore.get('identityKey', null)) as libsignal.IdentityKeyPair;
    const bobPreKeyBundle = {
        identityKey: bobIdentity.pubKey,
        registrationId: (await bobStore.get('registrationId', null)) as number,
        preKey: await libsignal.KeyHelper.generateLastResortPreKey(),
        signedPreKey: (await libsignal.KeyHelper.generateSignedPreKey(bobIdentity, 0)),
    };

    await aliceSessionBuilder.processPreKey(bobPreKeyBundle);

    return { aliceStore, bobStore };
}