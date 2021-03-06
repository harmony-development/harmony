package v1

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/harmony-development/hrpc/server"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func BindPB(obj interface{}, c echo.Context) error {
	buf, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	ct := c.Request().Header.Get("Content-Type")
	switch ct {
	case "application/hrpc", "application/octet-stream":
		if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
			return err
		}
	case "application/hrpc-json":
		if err = protojson.Unmarshal(buf, obj.(proto.Message)); err != nil {
			return err
		}
	}

	return nil
}

var Mediaproxyᐳv1ᐳmediaproxy *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x1emediaproxy/v1/mediaproxy.proto\x12\x16protocol.mediaproxy.v1\x1a\x1bharmonytypes/v1/types.proto\"\xaa\x01\n\fSiteMetadata\x12\x1d\n\nsite_title\x18\x01 \x01(\tR\tsiteTitle\x12\x1d\n\npage_title\x18\x02 \x01(\tR\tpageTitle\x12\x12\n\x04kind\x18\x03 \x01(\tR\x04kind\x12 \n\vdescription\x18\x04 \x01(\tR\vdescription\x12\x10\n\x03url\x18\x05 \x01(\tR\x03url\x12\x14\n\x05image\x18\x06 \x01(\tR\x05image\"G\n\rMediaMetadata\x12\x1a\n\bmimetype\x18\x01 \x01(\tR\bmimetype\x12\x1a\n\bfilename\x18\x02 \x01(\tR\bfilename\",\n\x18FetchLinkMetadataRequest\x12\x10\n\x03url\x18\x01 \x01(\tR\x03url\"\xa8\x01\n\x19FetchLinkMetadataResponse\x12?\n\ais_site\x18\x01 \x01(\v2$.protocol.mediaproxy.v1.SiteMetadataH\x00R\x06isSite\x12B\n\bis_media\x18\x02 \x01(\v2%.protocol.mediaproxy.v1.MediaMetadataH\x00R\aisMediaB\x06\n\x04data\"&\n\x12InstantViewRequest\x12\x10\n\x03url\x18\x01 \x01(\tR\x03url\"\x8c\x01\n\x13InstantViewResponse\x12@\n\bmetadata\x18\x01 \x01(\v2$.protocol.mediaproxy.v1.SiteMetadataR\bmetadata\x12\x18\n\acontent\x18\x02 \x01(\tR\acontent\x12\x19\n\bis_valid\x18\x03 \x01(\bR\aisValid\"B\n\x16CanInstantViewResponse\x12(\n\x10can_instant_view\x18\x01 \x01(\bR\x0ecanInstantView2\xf8\x02\n\x11MediaProxyService\x12\u007f\n\x11FetchLinkMetadata\x120.protocol.mediaproxy.v1.FetchLinkMetadataRequest\x1a1.protocol.mediaproxy.v1.FetchLinkMetadataResponse\"\x05\x9aD\x02\b\x01\x12m\n\vInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a+.protocol.mediaproxy.v1.InstantViewResponse\"\x05\x9aD\x02\b\x01\x12s\n\x0eCanInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a..protocol.mediaproxy.v1.CanInstantViewResponse\"\x05\x9aD\x02\b\x01B9Z7github.com/harmony-development/legato/gen/mediaproxy/v1J\x8d\v\n\x06\x12\x04\x00\x004\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\b\n\x01\x02\x12\x03\x02\x00\x1f\n\t\n\x02\x03\x00\x12\x03\x03\x00%\n\b\n\x01\b\x12\x03\x05\x00N\n\t\n\x02\b\v\x12\x03\x05\x00N\n\n\n\x02\x04\x00\x12\x04\a\x00\x0e\x01\n\n\n\x03\x04\x00\x01\x12\x03\a\b\x14\n\v\n\x04\x04\x00\x02\x00\x12\x03\b\x04\x1a\n\f\n\x05\x04\x00\x02\x00\x05\x12\x03\b\x04\n\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\b\v\x15\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\b\x18\x19\n\v\n\x04\x04\x00\x02\x01\x12\x03\t\x04\x1a\n\f\n\x05\x04\x00\x02\x01\x05\x12\x03\t\x04\n\n\f\n\x05\x04\x00\x02\x01\x01\x12\x03\t\v\x15\n\f\n\x05\x04\x00\x02\x01\x03\x12\x03\t\x18\x19\n\v\n\x04\x04\x00\x02\x02\x12\x03\n\x04\x14\n\f\n\x05\x04\x00\x02\x02\x05\x12\x03\n\x04\n\n\f\n\x05\x04\x00\x02\x02\x01\x12\x03\n\v\x0f\n\f\n\x05\x04\x00\x02\x02\x03\x12\x03\n\x12\x13\n\v\n\x04\x04\x00\x02\x03\x12\x03\v\x04\x1b\n\f\n\x05\x04\x00\x02\x03\x05\x12\x03\v\x04\n\n\f\n\x05\x04\x00\x02\x03\x01\x12\x03\v\v\x16\n\f\n\x05\x04\x00\x02\x03\x03\x12\x03\v\x19\x1a\n\v\n\x04\x04\x00\x02\x04\x12\x03\f\x04\x13\n\f\n\x05\x04\x00\x02\x04\x05\x12\x03\f\x04\n\n\f\n\x05\x04\x00\x02\x04\x01\x12\x03\f\v\x0e\n\f\n\x05\x04\x00\x02\x04\x03\x12\x03\f\x11\x12\n\v\n\x04\x04\x00\x02\x05\x12\x03\r\x04\x15\n\f\n\x05\x04\x00\x02\x05\x05\x12\x03\r\x04\n\n\f\n\x05\x04\x00\x02\x05\x01\x12\x03\r\v\x10\n\f\n\x05\x04\x00\x02\x05\x03\x12\x03\r\x13\x14\n\n\n\x02\x04\x01\x12\x04\x0f\x00\x12\x01\n\n\n\x03\x04\x01\x01\x12\x03\x0f\b\x15\n\v\n\x04\x04\x01\x02\x00\x12\x03\x10\x04\x18\n\f\n\x05\x04\x01\x02\x00\x05\x12\x03\x10\x04\n\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\x10\v\x13\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\x10\x16\x17\n\v\n\x04\x04\x01\x02\x01\x12\x03\x11\x04\x18\n\f\n\x05\x04\x01\x02\x01\x05\x12\x03\x11\x04\n\n\f\n\x05\x04\x01\x02\x01\x01\x12\x03\x11\v\x13\n\f\n\x05\x04\x01\x02\x01\x03\x12\x03\x11\x16\x17\n\n\n\x02\x04\x02\x12\x04\x14\x00\x16\x01\n\n\n\x03\x04\x02\x01\x12\x03\x14\b \n\v\n\x04\x04\x02\x02\x00\x12\x03\x15\x04\x13\n\f\n\x05\x04\x02\x02\x00\x05\x12\x03\x15\x04\n\n\f\n\x05\x04\x02\x02\x00\x01\x12\x03\x15\v\x0e\n\f\n\x05\x04\x02\x02\x00\x03\x12\x03\x15\x11\x12\n\n\n\x02\x04\x03\x12\x04\x17\x00\x1c\x01\n\n\n\x03\x04\x03\x01\x12\x03\x17\b!\n\f\n\x04\x04\x03\b\x00\x12\x04\x18\x04\x1b\x05\n\f\n\x05\x04\x03\b\x00\x01\x12\x03\x18\n\x0e\n\v\n\x04\x04\x03\x02\x00\x12\x03\x19\b!\n\f\n\x05\x04\x03\x02\x00\x06\x12\x03\x19\b\x14\n\f\n\x05\x04\x03\x02\x00\x01\x12\x03\x19\x15\x1c\n\f\n\x05\x04\x03\x02\x00\x03\x12\x03\x19\x1f \n\v\n\x04\x04\x03\x02\x01\x12\x03\x1a\b#\n\f\n\x05\x04\x03\x02\x01\x06\x12\x03\x1a\b\x15\n\f\n\x05\x04\x03\x02\x01\x01\x12\x03\x1a\x16\x1e\n\f\n\x05\x04\x03\x02\x01\x03\x12\x03\x1a!\"\n\n\n\x02\x04\x04\x12\x04\x1e\x00 \x01\n\n\n\x03\x04\x04\x01\x12\x03\x1e\b\x1a\n\v\n\x04\x04\x04\x02\x00\x12\x03\x1f\x04\x13\n\f\n\x05\x04\x04\x02\x00\x05\x12\x03\x1f\x04\n\n\f\n\x05\x04\x04\x02\x00\x01\x12\x03\x1f\v\x0e\n\f\n\x05\x04\x04\x02\x00\x03\x12\x03\x1f\x11\x12\n\n\n\x02\x04\x05\x12\x04!\x00%\x01\n\n\n\x03\x04\x05\x01\x12\x03!\b\x1b\n\v\n\x04\x04\x05\x02\x00\x12\x03\"\x04\x1e\n\f\n\x05\x04\x05\x02\x00\x06\x12\x03\"\x04\x10\n\f\n\x05\x04\x05\x02\x00\x01\x12\x03\"\x11\x19\n\f\n\x05\x04\x05\x02\x00\x03\x12\x03\"\x1c\x1d\n\v\n\x04\x04\x05\x02\x01\x12\x03#\x04\x17\n\f\n\x05\x04\x05\x02\x01\x05\x12\x03#\x04\n\n\f\n\x05\x04\x05\x02\x01\x01\x12\x03#\v\x12\n\f\n\x05\x04\x05\x02\x01\x03\x12\x03#\x15\x16\n\v\n\x04\x04\x05\x02\x02\x12\x03$\x04\x16\n\f\n\x05\x04\x05\x02\x02\x05\x12\x03$\x04\b\n\f\n\x05\x04\x05\x02\x02\x01\x12\x03$\t\x11\n\f\n\x05\x04\x05\x02\x02\x03\x12\x03$\x14\x15\n\n\n\x02\x04\x06\x12\x04&\x00(\x01\n\n\n\x03\x04\x06\x01\x12\x03&\b\x1e\n\v\n\x04\x04\x06\x02\x00\x12\x03'\x04\x1e\n\f\n\x05\x04\x06\x02\x00\x05\x12\x03'\x04\b\n\f\n\x05\x04\x06\x02\x00\x01\x12\x03'\t\x19\n\f\n\x05\x04\x06\x02\x00\x03\x12\x03'\x1c\x1d\n\n\n\x02\x06\x00\x12\x04*\x004\x01\n\n\n\x03\x06\x00\x01\x12\x03*\b\x19\n\f\n\x04\x06\x00\x02\x00\x12\x04+\x04-\x05\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03+\b\x19\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03+\x1a2\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03+=V\n\f\n\x05\x06\x00\x02\x00\x04\x12\x03,\bI\n\x0f\n\b\x06\x00\x02\x00\x04\xc3\b\x01\x12\x03,\bI\n\f\n\x04\x06\x00\x02\x01\x12\x04.\x040\x05\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03.\b\x13\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03.\x14&\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03.1D\n\f\n\x05\x06\x00\x02\x01\x04\x12\x03/\bI\n\x0f\n\b\x06\x00\x02\x01\x04\xc3\b\x01\x12\x03/\bI\n\f\n\x04\x06\x00\x02\x02\x12\x041\x043\x05\n\f\n\x05\x06\x00\x02\x02\x01\x12\x031\b\x16\n\f\n\x05\x06\x00\x02\x02\x02\x12\x031\x17)\n\f\n\x05\x06\x00\x02\x02\x03\x12\x0314J\n\f\n\x05\x06\x00\x02\x02\x04\x12\x032\bI\n\x0f\n\b\x06\x00\x02\x02\x04\xc3\b\x01\x12\x032\bIb\x06proto3")

	err := proto.Unmarshal(data, Mediaproxyᐳv1ᐳmediaproxy)
	if err != nil {
		panic(err)
	}
}

var MediaProxyServiceData *descriptorpb.ServiceDescriptorProto = new(descriptorpb.ServiceDescriptorProto)

func init() {
	data := []byte("\n\x11MediaProxyService\x12\u007f\n\x11FetchLinkMetadata\x120.protocol.mediaproxy.v1.FetchLinkMetadataRequest\x1a1.protocol.mediaproxy.v1.FetchLinkMetadataResponse\"\x05\x9aD\x02\b\x01\x12m\n\vInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a+.protocol.mediaproxy.v1.InstantViewResponse\"\x05\x9aD\x02\b\x01\x12s\n\x0eCanInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a..protocol.mediaproxy.v1.CanInstantViewResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, MediaProxyServiceData)
	if err != nil {
		panic(err)
	}
}

type MediaProxyServiceServer interface {
	FetchLinkMetadata(ctx echo.Context, r *FetchLinkMetadataRequest) (resp *FetchLinkMetadataResponse, err error)

	InstantView(ctx echo.Context, r *InstantViewRequest) (resp *InstantViewResponse, err error)

	CanInstantView(ctx echo.Context, r *InstantViewRequest) (resp *CanInstantViewResponse, err error)
}

var MediaProxyServiceServerFetchLinkMetadataData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x11FetchLinkMetadata\x120.protocol.mediaproxy.v1.FetchLinkMetadataRequest\x1a1.protocol.mediaproxy.v1.FetchLinkMetadataResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, MediaProxyServiceServerFetchLinkMetadataData)
	if err != nil {
		panic(err)
	}
}

var MediaProxyServiceServerInstantViewData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\vInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a+.protocol.mediaproxy.v1.InstantViewResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, MediaProxyServiceServerInstantViewData)
	if err != nil {
		panic(err)
	}
}

var MediaProxyServiceServerCanInstantViewData *descriptorpb.MethodDescriptorProto = new(descriptorpb.MethodDescriptorProto)

func init() {
	data := []byte("\n\x0eCanInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a..protocol.mediaproxy.v1.CanInstantViewResponse\"\x05\x9aD\x02\b\x01")

	err := proto.Unmarshal(data, MediaProxyServiceServerCanInstantViewData)
	if err != nil {
		panic(err)
	}
}

type MediaProxyServiceHandler struct {
	Server       MediaProxyServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	UnaryPre     server.HandlerTransformer
	upgrader     websocket.Upgrader
}

func NewMediaProxyServiceHandler(s MediaProxyServiceServer) *MediaProxyServiceHandler {
	return &MediaProxyServiceHandler{
		Server: s,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
	}
}

func (h *MediaProxyServiceHandler) SetUnaryPre(s server.HandlerTransformer) {
	h.UnaryPre = s
}

func (h *MediaProxyServiceHandler) Routes() map[string]echo.HandlerFunc {
	return map[string]echo.HandlerFunc{

		"/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata": h.FetchLinkMetadataHandler,

		"/protocol.mediaproxy.v1.MediaProxyService/InstantView": h.InstantViewHandler,

		"/protocol.mediaproxy.v1.MediaProxyService/CanInstantView": h.CanInstantViewHandler,
	}
}

func (h *MediaProxyServiceHandler) FetchLinkMetadataHandler(c echo.Context) error {

	requestProto := new(FetchLinkMetadataRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.FetchLinkMetadata(c, req.(*FetchLinkMetadataRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(MediaProxyServiceServerFetchLinkMetadataData, MediaProxyServiceData, Mediaproxyᐳv1ᐳmediaproxy, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *MediaProxyServiceHandler) InstantViewHandler(c echo.Context) error {

	requestProto := new(InstantViewRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.InstantView(c, req.(*InstantViewRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(MediaProxyServiceServerInstantViewData, MediaProxyServiceData, Mediaproxyᐳv1ᐳmediaproxy, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}

func (h *MediaProxyServiceHandler) CanInstantViewHandler(c echo.Context) error {

	requestProto := new(InstantViewRequest)
	if err := BindPB(requestProto, c); err != nil {
		return err
	}

	invoker := func(c echo.Context, req proto.Message) (proto.Message, error) {
		return h.Server.CanInstantView(c, req.(*InstantViewRequest))
	}

	if h.UnaryPre != nil {
		invoker = h.UnaryPre(MediaProxyServiceServerCanInstantViewData, MediaProxyServiceData, Mediaproxyᐳv1ᐳmediaproxy, invoker)
	}

	res, err := invoker(c, requestProto)
	if err != nil {
		return err
	}
	var response []byte

	ct := c.Request().Header.Get("Content-Type")

	switch ct {
	case "application/hrpc-json":
		response, err = protojson.Marshal(res)
	default:
		response, err = proto.Marshal(res)
	}

	if err != nil {
		return err
	}

	if ct == "application/hrpc-json" {
		return c.Blob(http.StatusOK, "application/hrpc-json", response)
	}
	return c.Blob(http.StatusOK, "application/hrpc", response)

}
