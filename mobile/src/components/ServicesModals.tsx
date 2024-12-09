import React, { useState } from 'react';
import { View, Text, Modal, StyleSheet, TouchableOpacity } from 'react-native';
import ModalDropdown from 'react-native-modal-dropdown';
import { Action, Reaction, Service, ServicesModalsProps } from '../types';
import { Button } from 'react-native-paper';

export function ServicesModals({
  modalVisible,
  setModalVisible,
  services,
  isAction,
  setActionOrReaction,
}: ServicesModalsProps) {
  const [selectedService, setSelectedService] = useState<string | null>(null);
  const [selectedActionOrReaction, setSelectedActionOrReaction] = useState<
    string | null
  >(null);
  const [id, setId] = useState<number | null>(null);

  const serviceOptions =
    services?.server?.services?.map((service: Service) => service.name) || [];

  const selectedServiceData = selectedService
    ? services?.server?.services?.find(
        service => service.name === selectedService,
      )
    : null;

  const actionOrReactionOptions = isAction
    ? selectedServiceData?.actions?.map((action: Action) => ({
        name: action.name,
        id: action.action_id,
      })) || []
    : selectedServiceData?.reactions?.map((reaction: Reaction) => ({
        name: reaction.name,
        id: reaction.reaction_id,
      })) || [];

  const handleSave = () => {
    if (selectedActionOrReaction !== null) {
      if (setActionOrReaction && id) {
        setActionOrReaction(1);
      } else {
        console.error('setActionOrReaction is not defined');
      }
    }
    setModalVisible(false);
    setSelectedService(null);
    setSelectedActionOrReaction(null);
    setId(null);
  };

  return (
    <Modal
      animationType="fade"
      transparent={true}
      visible={modalVisible}
      onRequestClose={() => {
        setModalVisible(!modalVisible);
      }}>
      <View style={styles.modalOverlay}>
        <View style={styles.modalContainer}>
          <Text style={styles.modalText}>Make Your Workflows</Text>
          <View style={styles.container}>
            <Text style={styles.label}>Select a service:</Text>
            <ModalDropdown
              options={serviceOptions}
              key={serviceOptions.length}
              defaultValue="Select a service"
              onSelect={(index, value) => {
                setSelectedService(value);
                setSelectedActionOrReaction(null);
              }}
              style={styles.dropdown}
              textStyle={styles.dropdownText}
              dropdownTextStyle={styles.dropdownItemText}
              renderRow={(option, index, isSelected) => (
                <Text
                  key={index}
                  style={isSelected ? styles.selectedItem : styles.item}>
                  {option}
                </Text>
              )}
            />
            {selectedService && (
              <>
                <Text style={styles.selectedValue}>
                  Select an {isAction ? 'action' : 'reaction'} for{' '}
                  {selectedService}
                </Text>
                <ModalDropdown
                  options={actionOrReactionOptions.map(option => option.name)}
                  defaultValue={`Select an ${isAction ? 'action' : 'reaction'}`}
                  onSelect={(index, value) => {
                    setSelectedActionOrReaction(value);
                    const selectedOption = actionOrReactionOptions.find(
                      option => option.name === value,
                    );
                    if (selectedOption) {
                      setId(selectedOption.id);
                    }
                  }}
                  key={actionOrReactionOptions.length}
                  style={styles.dropdown}
                  textStyle={styles.dropdownText}
                  dropdownTextStyle={styles.dropdownItemText}
                  renderRow={(option, index, isSelected) => (
                    <Text
                      key={index}
                      style={isSelected ? styles.selectedItem : styles.item}>
                      {option}
                    </Text>
                  )}
                />
                {selectedActionOrReaction && (
                  <Text style={styles.selectedValue}>
                    Selected {isAction ? 'action' : 'reaction'}:{' '}
                    {selectedActionOrReaction}
                  </Text>
                )}
              </>
            )}
            <View style={styles.buttonContainer}>
              <Button style={styles.saveButton} onPress={handleSave}>
                <Text style={{ color: 'white', fontSize: 16 }}>Save</Text>
              </Button>
              <Button
                style={styles.cancelButton}
                onPress={() => {
                  setModalVisible(false);
                  setSelectedService(null);
                  setSelectedActionOrReaction(null);
                  setId(null);
                }}>
                <Text style={{ color: 'black', fontSize: 16 }}>Cancel</Text>
              </Button>
            </View>
          </View>
        </View>
      </View>
    </Modal>
  );
}

const styles = StyleSheet.create({
  actionReactionContainer: {
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
    gap: 10,
    marginVertical: 20,
  },
  modalOverlay: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'rgba(0, 0, 0, 0.5)',
  },
  modalContainer: {
    width: '90%',
    padding: 20,
    backgroundColor: 'white',
    borderRadius: 10,
    alignItems: 'center',
  },
  modalText: {
    fontSize: 18,
    marginBottom: 20,
  },
  button: {
    padding: 10,
    backgroundColor: '#ccc',
    borderRadius: 5,
    margin: 5,
  },
  selectedButton: {
    backgroundColor: '#007BFF',
  },
  buttonText: {
    color: 'white',
    fontSize: 16,
  },
  label: {
    fontSize: 18,
    marginBottom: 10,
  },
  dropdown: {
    width: 'auto',
    height: 'auto',
    borderWidth: 1,
    borderColor: '#ccc',
    borderRadius: 5,
    justifyContent: 'center',
    padding: 10,
  },
  dropdownText: {
    fontSize: 16,
  },
  dropdownItemText: {
    fontSize: 16,
    padding: 10,
  },
  selectedValue: {
    marginTop: 20,
    fontSize: 18,
  },
  container: {
    width: '100%',
    alignItems: 'center',
  },
  item: {
    padding: 10,
    fontSize: 16,
  },
  selectedItem: {
    padding: 10,
    fontSize: 16,
    backgroundColor: '#ddd',
  },
  buttonContainer: {
    marginTop: 20,
    flexDirection: 'row',
    justifyContent: 'space-between',
    width: '100%',
  },
  saveButton: {
    width: '48%',
    backgroundColor: 'red',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 20,
  },
  cancelButton: {
    width: '48%',
    backgroundColor: '#E8E9E9',
    alignItems: 'center',
    justifyContent: 'center',
  },
});