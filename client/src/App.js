import axios from 'axios';
import React, {  useEffect, useCallback, useState } from 'react';
import { client } from "./index";
import gql from 'graphql-tag';

const BACKGROUND_COLORS = [
    'red',
    'blue',
    'yellow',
    'green',
    'gray',
    'orange',
    'beige',
    'aquamarine'
];

const App = () => {
    const [bgColor, setBgColor] = useState('');
    const [todos, setTodos] = useState([]);

    const fetchTodos = useCallback(() => {
        client.query({
            query: gql(`
                query {
                  todos {
                    id
                    name
                  }
                }
              `)
        })
            .then((results) => {
                const result = results.data;

                setTodos(result.todos);
            })
            .catch((error) => {
                console.error(error);
            });
    })

    const createTodo = useCallback((e) => {
        e.preventDefault();

        const doc = gql(`
    mutation createTodo ($input: TodoInput!) {
          createTodo (input: $input) {
              name
          }
      }`);
        client.mutate({
            mutation: doc,
            variables: {
                input: {
                    name: 'testName',
                },
            }})
            .then((results) => {
                const result = results.data;

                fetchTodos();
            })
            .catch((error) => {
                console.error(error);
            });
    }, []);

    const setSettings = useCallback((e) => {
        e.preventDefault();
        const { value } = e.target;

        axios.post('api/rest/settings', { bgColor: value }).then((response) => {
            setBgColor(value);
        });
    }, []);

    useEffect(() => {
        axios.get('api/rest/settings').then((response) => {
            console.log(response)
            const {bgColor = ''} = response.data.result;

            setBgColor(bgColor);
        });

        fetchTodos();
    });

    return (
        <div>
            <form>
                <div>
                    <button onClick={createTodo}>createTodo</button>
                </div>
                {BACKGROUND_COLORS.map((value, index) => {
                    let checked = false;

                    if (bgColor === value) {
                        checked = true;
                    }

                    return (
                        <span key={index}>
                            <label htmlFor={value}>{ value }</label>
                            <input id={value} name="bg_color" type="radio" value={value} checked={checked} onChange={setSettings} />
                        </span>
                    );
                })}
                <div>
                    선택된 배경색:
                    <span
                        style={{
                            width: 10,
                            height: 10,
                            backgroundColor: bgColor,
                            display: 'inline-block'
                        }}>
                        &nbsp;
                    </span>
                </div>
            </form>
            {todos.map((v) => {
                return (
                    <div>{v.name}_{v.id}</div>
                )
            })}
        </div>
    );
};

export default App;
