import { atom } from "recoil";
import * as WebSocket from "websocket";

export const websocketAtom = atom<WebSocket.w3cwebsocket>({
  key: "websocket",
  default: new WebSocket.w3cwebsocket("ws://localhost:80/ws"),
});
