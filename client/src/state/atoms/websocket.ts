import { atom } from "recoil";
import * as WebSocket from "websocket";

export const websocketAtom = atom<WebSocket.w3cwebsocket>({
  key: "websocket",
  default: undefined,
});
