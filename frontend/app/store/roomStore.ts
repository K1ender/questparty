import type { Response } from "~/types/response";

export const useRoomStore = defineStore("room", {
  state: () => ({
    connected: false,
    roomId: "",
    username: "",
  }),
  actions: {
    async joinRoom(username: string) {
      const {
        public: { apiBase },
      } = useRuntimeConfig();
      const resp = await $fetch<Response<{ roomId: string }>>(
        `${apiBase}/api/create`,
        {
          method: "POST",
          body: {
            username,
          },
        }
      );

      this.connected = true;
      this.roomId = resp.data.roomId;
      this.username = username;

      console.log(resp.data.roomId);
    },
  },
  getters: {
    isConnected: (state) => state.connected,
  },
});
