import React, {useState} from 'react';
import {Tooltip, Tag, List, Button, Popconfirm, Switch} from 'antd';
import {CloseOutlined, CheckOutlined} from '@ant-design/icons';

const Todo = ({todo, onTodoRemoval, onTodoToggle}) => {

    return (
        <List.Item
        actions = {[
            <Tooltip
            title={todo.completed ? 'Mark as uncompleted':'Mark as completed'}>
                <Switch 
                    checkedChildren={<CheckOutlined/>}
                    unCheckedChildren = {<CloseOutlined/>}
                    onChange={()=> onTodoToggle(todo)}
                    defaultChecked={todo.completed}
                />
            </Tooltip>,
            <Popconfirm
            tilte={'Are you sure want to delete?'}
            onConfirm={() =>{
                onTodoRemoval(todo);
            }}>
                <Button className="remove-todo-button" type="primary" danger>
                 X
                </Button>
            </Popconfirm>
        ]}
        className = "list-item"
        key={todo.id}
        >
            <div className="todo-item">
                <Tag color={todo.completed ? 'cyan' : 'red'} className="todo-tag">
                    {todo.title}
                </Tag>
            </div>

         </List.Item>
    )
}

export default Todo;