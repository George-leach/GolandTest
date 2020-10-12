const port = process.env.PORT
const express = require('express')
const app = express()


//static
app.use(express.static('public'))
app.use('/css',express.static(__dirname+ 'public/css'))
app.use('/js',express.static(__dirname+ 'public/js'))
app.use('/img',express.static(__dirname+ 'public/img'))

app.get('',(req,res)=>{
    res.sendFile(__dirname+'/views/index.html')
})

app.list(port,() =>console.info(`listening on port ${port}`))