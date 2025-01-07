import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { Button } from 'react-native-paper';
import { DetailsModals } from './DetailsModals';
import { globalStyles } from '../styles/global_style';

type Workflow = {
  name: string;
  action_id: number;
  reaction_id: number;
  is_active: boolean;
  created_at: string;
};

type WorkflowTableProps = {
  workflows: Workflow[];
  detailsModalVisible: boolean;
  setDetailsModalVisible: (detailsModalVisible: boolean) => void;
};

export function WorkflowTable({
  workflows,
  setDetailsModalVisible,
  detailsModalVisible,
}: WorkflowTableProps) {
  return (
    <>
      <View style={styles.container}>
        <View style={styles.headerRow}>
          <Text style={styles.headerCell}>Name</Text>
          <Text style={styles.headerCell}>Is Active</Text>
        </View>
        {workflows !== null &&
          workflows.map((workflow, index) => (
            <View key={index} style={styles.row}>
              <Text style={styles.cell}>{workflow.name}</Text>
              <Text style={styles.cell}>
                {workflow.is_active ? 'Yes' : 'No'}
              </Text>
              <Button
                onPress={() => setDetailsModalVisible(!detailsModalVisible)}>
                <Text>details</Text>
              </Button>
            </View>
          ))}
      </View>
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    marginTop: 10,
    width: '100%',
    borderColor: '#ccc',
    borderRadius: 5,
  },
  headerRow: {
    flexDirection: 'row',
    backgroundColor: '#FFFFFF',
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
