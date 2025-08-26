<script lang="ts" setup>
import { useRoomStore } from '~/store/roomStore';

const isLoading = ref(false);
const roomStore = useRoomStore();


async function createRoom() {
    isLoading.value = true;
    try {
        await roomStore.joinRoom("test");
        navigateTo(`/room/${roomStore.roomId}`);
    } catch (e) {
        console.error(e);
    } finally {
        isLoading.value = false;
    }
}
</script>

<template>
    <div class="container">
        <div class="card">
            <h1>Quest Party</h1>
            <button :disabled="isLoading" @click="createRoom">{{
                isLoading ? "Creating..." : "Create Room"
            }}</button>
        </div>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    justify-content: center;
    align-items: center;
}

.card {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

button {
    background-color: hsl(216, 100%, 50%);
    border: 0;
    color: white;

    font-size: 1.4rem;
    width: 100%;
    height: 4rem;
    border-radius: 0.5rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.1s ease-out;
}

button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

button:hover {
    background-color: hsl(216, 100%, 45%);
}
</style>