import { useEffect, useState } from 'react';
import { Button, ScrollView, StyleSheet, Text, View } from 'react-native';
import { Modal } from 'react-native';
import { getWorkflows } from '../service/workflows';
import { PullRequestComment } from '../types';

interface DetailsModalsProps {
  modalVisible: boolean;
  setModalVisible: (modalVisible: boolean) => void;
  workflows: PullRequestComment[];
}

export function DetailsModals({
  modalVisible,
  setModalVisible,
  workflows,
}: DetailsModalsProps) {
  return (
    <ScrollView style={styles.container}>
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
          {workflows &&
            workflows.map((workflows, index) => (
              <View key={index} style={styles.row}>
                <Text style={styles.cell}>{workflows.body}</Text>
                <Text style={styles.cell}>{workflows.pull_request_url}</Text>
              </View>
            ))}
        </View>
      </Modal>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    margin: 10,
    borderWidth: 1,
    borderColor: '#ccc',
    borderRadius: 5,
    backgroundColor: 'white',
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
