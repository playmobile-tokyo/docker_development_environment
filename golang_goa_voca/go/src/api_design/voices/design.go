package design                                 // The convention consists of naming the design
                                                   // package "design"
import (
        . "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("vocatuku", func() {                     // API defines the microservice endpoint and
        Title("The virtual wine vocatuku")           // other global properties. There should be one
        Description("A simple goa service")        // and exactly one API definition appearing in
        Scheme("http")                             // the design.
        Host("localhost:5000")
})

var _ = Resource("voice", func() {                // Resources group related API endpoints
        BasePath("/voices")                       // together. They map to REST resources for REST
        DefaultMedia(VoiceMedia)                  // services.

        Action("show", func() {                    // Actions define a single API endpoint together
                Description("Get voice by id")    // with its path, parameters (both path
                Routing(GET("/:voiceID"))         // parameters and querystring values) and payload
                Params(func() {                    // (shape of the request body).
                        Param("voiceID", Integer, "Voice ID")
                })
                Response(OK)                       // Responses define the shape and status code
                Response(NotFound)                 // of HTTP responses.
        })
})

var VoiceMedia = MediaType("application/vnd.goa.voice+json", func() {
        Description("A voice of you")
        Attributes(func() {                         // Attributes define the media type shape.
                Attribute("id", Integer, "Unique voice ID")
                Attribute("href", String, "API href for making requests on the voice")
                Attribute("name", String, "Name of wine")
                Required("id", "href", "name")
        })
        View("default", func() {                    // View defines a rendering of the media type.
                Attribute("id")                     // Media types may have multiple views and must
                Attribute("href")                   // have a "default" view.
                Attribute("name")
        })
})
