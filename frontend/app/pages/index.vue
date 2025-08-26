<script lang="ts" setup>
async function createRoom() {
    // await fetch("http://localhost:8080/api/room", {
    //     method: "POST",
    // });
}
</script>

<template>
    <div>
        <h1>Quest Party</h1>
        <button>Create room</button>
    </div>
</template>

<style scoped>
div {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    height: 100vh;
    justify-content: center;
    align-items: center;
}

button {
    background-color: hsl(216, 100%, 50%);
    border: 0;
    color: white;
    padding: 1rem 2rem;
    border-radius: 0.5rem;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.1s ease-out;
}

button:hover {
    background-color: hsl(216, 100%, 45%);
}
</style>