import { atom, selector } from "recoil";
import * as WebSocket from "websocket";

const connect = (): Promise<WebSocket.w3cwebsocket> => {
  return new Promise((resolve, reject) => {
    const port = import.meta.env.VITE_WS_PORT;
    const url = "ws://localhost:" + port + "/ws";
    const socket = new WebSocket.w3cwebsocket(url);

    socket.onopen = () => {
      console.log("connected", port);
      resolve(socket);
    };
    socket.onclose = () => {
      console.log("reconnecting...");
      connect();
    };
    socket.onerror = (err) => {
      console.log("connection error:", err);
      reject(err);
    };
  });
};

const connectWebsocketSelector = selector({
  key: "connectWebsocket",
  get: async (): Promise<WebSocket.w3cwebsocket> => {
    return await connect();
  },
});

export const websocketAtom = atom<WebSocket.w3cwebsocket>({
  key: "websocket",
  default: connectWebsocketSelector,
});
