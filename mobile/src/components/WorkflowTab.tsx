import React from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { Button } from 'react-native-paper';
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
    <ScrollView horizontal>
      <View style={styles.container}>
        <View style={[styles.headerRow, globalStyles.secondaryDark]}>
          <Text style={[styles.headerCell, globalStyles.textBlack]}>Name</Text>
          <Text style={[styles.headerCell, globalStyles.textBlack]}>
            Action
          </Text>
          <Text style={[styles.headerCell, globalStyles.textBlack]}>
            Reaction
          </Text>
          <Text style={[styles.headerCell, globalStyles.textBlack]}>
            Is Active
          </Text>
          <Text style={[styles.headerCell, globalStyles.textBlack]}>
            Details
          </Text>
        </View>
        {workflows &&
          workflows.map((workflow, index) => (
            <View key={index} style={[styles.row, globalStyles.secondaryDark]}>
              <View style={styles.cell}>
                <Text style={styles.text}>{workflow.name}</Text>
              </View>
              <View style={styles.cell}>
                <Text style={styles.text}>{workflow.action_id}</Text>
              </View>
              <View style={styles.cell}>
                <Text style={styles.text}>{workflow.reaction_id}</Text>
              </View>
              <View style={styles.cell}>
                <Text style={styles.text}>
                  {workflow.is_active ? 'Yes' : 'No'}
                </Text>
              </View>
              <View style={styles.cell}>
                <Button
                  mode="contained"
                  onPress={() => setDetailsModalVisible(!detailsModalVisible)}>
                  <Text style={styles.buttonText}>details</Text>
                </Button>
              </View>
            </View>
          ))}
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    padding: 10,
    width: '100%',
    borderColor: '#ccc',
    borderRadius: 5,
  },
  headerRow: {
    alignContent: 'center',
    flexDirection: 'row',
    borderTopLeftRadius: 5,
    borderTopRightRadius: 5,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  headerCell: {
    width: 100,
    flex: 1,
    padding: 10,
    fontWeight: 'bold',
    textAlign: 'center',
    color: '#333',
  },
  row: {
    alignContent: 'center',
    flexDirection: 'row',
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
    justifyContent: 'center',
  },
  text: {
    color: '#F7FAFB',
    textAlign: 'center',
  },
  cell: {
    justifyContent: 'center',
    alignItems: 'center',
    width: 100,
  },
  buttonText: {
    color: '#fff',
  },
});
