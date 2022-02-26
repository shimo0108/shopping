import React from 'react';

import {
  Stack,
  Button,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalCloseButton,
  ModalBody,
  ModalFooter,
} from '@chakra-ui/react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  existingResutaurautName: string;
  newResutaurautName: string;
  onClickSubmit: () => void;
};

export const NewOrderConfirmModal = ({
  isOpen,
  onClose,
  existingResutaurautName,
  newResutaurautName,
  onClickSubmit,
}: Props) => (
  <Modal isOpen={isOpen} onClose={onClose}>
    <ModalOverlay />
    <ModalContent>
      <ModalHeader>新規注文を開始しますか？</ModalHeader>
      <ModalCloseButton />
      <ModalBody>
        <p>
          {`ご注文に ${existingResutaurautName} の商品が含まれています。
          新規の注文を開始して ${newResutaurautName} の商品を追加してください。`}
        </p>
      </ModalBody>
      <ModalFooter>
        <Stack direction='row' align='center'>
          <Button
            size='md'
            height='40px'
            width='200px'
            border='2px'
            onClick={onClickSubmit}
          >
            新規注文
          </Button>
        </Stack>
      </ModalFooter>
    </ModalContent>
  </Modal>
);
