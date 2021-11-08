import './App.css';
import { useState } from 'react'

const baseUrl = process.env.REACT_APP_API_URL

function App() {
  const [key, updateKey] = useState('')
  const [value, updateValue] = useState('')
  const [setKey, updateSetKey] = useState('')
  const [setValue, updateSetValue] = useState('')
  const [loading, setLoading]   = useState(true)

  const displayValue = () => {
    if (loading) {
      return ''
    } else {
      return value
    }
  }

  async function getValue(key) {
    const endpoint = `${baseUrl}/get?key=${key}`
    console.log(endpoint)
    const response = await fetch(endpoint, {})
    const result = await response.json()

    updateValue(result.value)
    setLoading(false)
  }

  async function setCacheValue(key, value) {
    const endpoint = `${baseUrl}/set`
    const body = JSON.stringify({ key, value })
    const response = await fetch(endpoint, { method: 'POST', body })
    const result = await response.json()

    console.log(result)
  }

  const handleGetSubmit = (event) => {
    event.preventDefault()
    getValue(key)
  }

  const handleSetSubmit = (event) => {
    event.preventDefault()
    setCacheValue(setKey, setValue)
  }

  const handleChange = (event) => {
    const { value } = event.target
    updateKey(value)
  }

  const handleSetKeyChange = (event) => {
    const { value } = event.target
    updateSetKey(value)
  }

  const handleSetValueChange = (event) => {
    const { value } = event.target
    updateSetValue(value)
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>My Cache</h1>

        <div>
          <h2>Get Value</h2>
          <form onSubmit={handleGetSubmit}>
            <label>
              Key:
              <input type="text" value={key} onChange={handleChange}></input>
            </label>
            <input type="submit" value="Submit" />
          </form>
          <h3>Value: {displayValue()} </h3>

          <h2>Set Value</h2>
          <form onSubmit={handleSetSubmit}>
            <label>
              Key:
              <input type="text" value={setKey} onChange={handleSetKeyChange}></input>
            </label>
            <label>
              Value:
              <input type="text" value={setValue} onChange={handleSetValueChange}></input>
            </label>
            <input type="submit" value="Submit" />
          </form>
        </div>
      </header>
    </div>
  )
}

export default App
