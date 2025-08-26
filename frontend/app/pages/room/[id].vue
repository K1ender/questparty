<script setup lang="ts">
import { useRoomStore } from '~/store/roomStore'

const roomStore = useRoomStore()
const route = useRoute()

useHead({
    title: "Quest Party - Room " + route.params.id,
})

onMounted(() => {
    if (!roomStore.connected) {
        navigateTo("/");
    }
})

function copyInviteLink() {
    navigator.clipboard.writeText(window.location.href);
}
</script>


<template>
    <div>
        <p>You are in room {{ route.params.id }}</p>
        <a @click="copyInviteLink">Copy invite link</a>
    </div>
</template>

<style lang="css" scoped>
a {
    cursor: pointer;
}
</style>