<!-- prettier-ignore-start -->
# API Docs
### These are temporary until Swagger is setup in this project. 
### Note you cannot directly query the api unless you have the cookies set from querying the login endpoint.
## /recommendations
artist_ids:
    type: array
    items:
        type: string
track_ids:
    type: array
    items:
        type: string
genres:
    type: array
    items: 
        type: string
recommendation_count:
    type: int32

## /trackfeatures
id:
    type: string

## /trackfeatures
id: 
    type: string

## /colors
query:
    type: string
search_type:
    type: string
    enum: [album, artist]
background_colors:
    type: array
    items:
        type: string (hex codes)
highlight_colors:
    type: array
    items:
        type: string (hex codes)



<!-- prettier-ignore-end -->
