import { Message } from "./../../models/message";
import { atom } from "recoil";

export const messageAtom = atom<Message>({
  key: "message",
  default: undefined,
});

export const messageListAtom = atom<Message[]>({
  key: "messageList",
  default: [],
});
