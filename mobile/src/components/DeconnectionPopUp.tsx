import React, { useState } from 'react';
import { View, Text, Modal, StyleSheet, TouchableOpacity } from 'react-native';
import { deleteToken, logoutServices } from '../service';
import { globalStyles } from '../styles/global_style';

interface DeconnectionPopUpProps {
  service: string;
  modalVisible: boolean;
  setModalVisible: (modalVisible: boolean) => void;
  token: string;
  serverIp: string;
  setNeedRefresh: (needRefresh: boolean) => void;
}

export function DeconnectionPopUp({
  service,
  modalVisible,
  setModalVisible,
  token,
  serverIp,
  setNeedRefresh
}: DeconnectionPopUpProps) {
  
  const handleDeconnection = async () => {
    await logoutServices(serverIp, service, token);
    deleteToken(service);
    setModalVisible(false);
    setNeedRefresh(true);
  };

  return (
    <View style={styles.container}>
      <Modal
        animationType="fade"
        transparent={true}
        visible={modalVisible}
        onRequestClose={() => {
          setModalVisible(!modalVisible);
        }}>
        <View style={styles.modalOverlay}>
          <View style={styles.modalContainer}>
            <Text style={styles.modalText}>
              Tu es déjà connecté à {service}, veux-tu te déconnecter ?
            </Text>
            <View style={styles.buttonContainer}>
              <TouchableOpacity
                style={[styles.button, globalStyles.buttonFormat]}
                onPress={handleDeconnection}>
                <Text style={{ color: 'white', fontSize: 16 }}>
                  Déconnecter
                </Text>
              </TouchableOpacity>
              <TouchableOpacity
                style={[styles.cancelButton, globalStyles.buttonFormat]}
                onPress={() => setModalVisible(false)}>
                <Text style={{ color: 'black', fontSize: 16 }}>Annuler</Text>
              </TouchableOpacity>
            </View>
          </View>
        </View>
      </Modal>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  modalOverlay: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
  },
  modalContainer: {
    width: '70%',
    padding: 20,
    backgroundColor: 'white',
    borderRadius: 20,
    alignItems: 'center',
  },
  modalText: {
    fontSize: 18,
    marginBottom: 20,
    textAlign: 'center',
  },
  buttonContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%',
  },
  button: {
    width: 'auto',
    backgroundColor: 'red',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 20,
  },
  cancelButton: {
    width: 'auto',
    backgroundColor: '#E8E9E9',
    alignItems: 'center',
    justifyContent: 'center',
  },
  buttonText: {
    color: 'white',
    fontSize: 16,
  },
});
