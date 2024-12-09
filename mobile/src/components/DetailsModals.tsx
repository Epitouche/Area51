import { useEffect } from 'react';
import { Button, StyleSheet, Text, View } from 'react-native';
import { Modal } from 'react-native';
import { getWorkflows } from '../service/workflows';

interface DetailsModalsProps {
  modalVisible: boolean;
  setModalVisible: (modalVisible: boolean) => void;
  serverIp: string;
  token: string;
}

export function DetailsModals({
  modalVisible,
  setModalVisible,
  serverIp,
  token,
}: DetailsModalsProps) {
  useEffect(() => {
    getWorkflows(serverIp, token);
  }), [modalVisible];
  return (
    <View style={styles.container}>
      <Modal
        animationType="fade"
        transparent={true}
        visible={modalVisible}
        onRequestClose={() => {
          setModalVisible(!modalVisible);
        }}>
        <View style={styles.container}>
          <View style={styles.headerRow}>
            <Text style={styles.headerCell}>Name</Text>
            <Text style={styles.headerCell}>Is Active</Text>
          </View>
          {/* {(
              <View key={index} style={styles.row}>
                <Text style={styles.cell}>{workflow.name}</Text>
                <Text style={styles.cell}>
                  {workflow.is_active ? 'Yes' : 'No'}
                </Text>
                <Button>
                  <Text>details</Text>
                </Button>
              </View>
            ))} */}
        </View>
      </Modal>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    margin: 10,
    borderWidth: 1,
    borderColor: '#ccc',
    borderRadius: 5,
  },
  headerRow: {
    flexDirection: 'row',
    backgroundColor: '#f8f8f8',
    borderTopLeftRadius: 5,
    borderTopRightRadius: 5,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  headerCell: {
    flex: 1,
    padding: 10,
    fontWeight: 'bold',
    textAlign: 'center',
  },
  row: {
    flexDirection: 'row',
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  cell: {
    flex: 1,
    padding: 10,
    textAlign: 'center',
  },
});
