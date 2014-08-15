go-color-analyzer
=================

A go library to extract the main color of an image. 

# Deploy on heroku

    git push heroku master
    
# Usage:
    
    GET $HOST/color?url=CGI_ESCAPED_URL
    
    RESPONSE>
    Encoding: text/plain
    Body: #HEX_COLOR
