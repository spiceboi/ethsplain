import './ethsplain.css'
import React, { useEffect } from 'react'
import cuid from 'cuid'

const sampleMatches = [
  {
    match: 'cool'
  },
  {
    match: 'cool'
  },
  {
    match: 'cool'
  },
  {
    match: 'cool'
  },
  {
    match: 'cool'
  },
  {
    match: 'cool'
  },
  {
    match: 'cool 2'
  }
]

const sampleHelptext = [
  {
    text: 'Hey there'
  },
  {
    text: 'Hey there'
  },
  {
    text: 'Hey there'
  },
  {
    text: 'Hey there'
  },
  {
    text: 'Hey there'
  },
  {
    text: 'Hey there'
  },
  {
    text: 'Hey there number 2'
  }
]

export default ({ helptext = sampleHelptext, matches = sampleMatches }) => {
  useEffect(() => {
    process.browser && window.init()
  }, [])

  return (
    <>
      <svg id='canvas' />
      <div id='command'>
        { matches.map((match, i) => <span key={cuid()} className='command0' helpref={`help-${i}`}><a>{match.match}</a></span>)}
      </div>
      <div style={{ height: sampleMatches.length * 10 + 'px' }} />
      <div id='help'>
        {helptext.map(({ text, id }, i) => (
          <pre key={cuid()} id={`help-${i}`} className='help-box help-synopsis'>{text}</pre>
        ))}
      </div>
    </>
  )
}
