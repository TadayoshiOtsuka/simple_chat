import { websocketAtom } from "./../state/atoms/websocket";
import { useRecoilState } from "recoil";
import { useEffect } from "react";
import * as WebSocket from "websocket";

export const useWebsocket = (): WebSocket.w3cwebsocket => {
  const [websocket, setWebsocket] = useRecoilState(websocketAtom);

  useEffect(() => {
    const socket = new WebSocket.w3cwebsocket("ws://localhost:80/ws");
    setWebsocket(socket);

    return () => {
      socket.close();
    };
  }, []);

  return websocket;
};
