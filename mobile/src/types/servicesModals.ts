import { AboutJson } from './aboutJson';

export interface ServicesModalsProps {
  modalVisible: boolean;
  setModalVisible: (modalVisible: boolean) => void;
  services: AboutJson | undefined;
  setActionOrReaction: (actionOrReaction: number) => void;
  isAction?: boolean;
}

export type ActionReaction = {
  id: number;
  name: string;
};
