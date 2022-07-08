package audits

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// AffectedCookie information about a cookie that is affected by an inspector
// issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-AffectedCookie
type AffectedCookie struct {
	Name   string `json:"name"` // The following three properties uniquely identify a cookie
	Path   string `json:"path"`
	Domain string `json:"domain"`
}

// AffectedRequest information about a request that is affected by an
// inspector issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-AffectedRequest
type AffectedRequest struct {
	RequestID network.RequestID `json:"requestId"` // The unique request id.
	URL       string            `json:"url,omitempty"`
}

// AffectedFrame information about the frame affected by an inspector issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-AffectedFrame
type AffectedFrame struct {
	FrameID cdp.FrameID `json:"frameId"`
}

// SameSiteCookieExclusionReason [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieExclusionReason
type SameSiteCookieExclusionReason string

// String returns the SameSiteCookieExclusionReason as string value.
func (t SameSiteCookieExclusionReason) String() string {
	return string(t)
}

// SameSiteCookieExclusionReason values.
const (
	SameSiteCookieExclusionReasonExcludeSameSiteUnspecifiedTreatedAsLax SameSiteCookieExclusionReason = "ExcludeSameSiteUnspecifiedTreatedAsLax"
	SameSiteCookieExclusionReasonExcludeSameSiteNoneInsecure            SameSiteCookieExclusionReason = "ExcludeSameSiteNoneInsecure"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SameSiteCookieExclusionReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SameSiteCookieExclusionReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SameSiteCookieExclusionReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SameSiteCookieExclusionReason(in.String()) {
	case SameSiteCookieExclusionReasonExcludeSameSiteUnspecifiedTreatedAsLax:
		*t = SameSiteCookieExclusionReasonExcludeSameSiteUnspecifiedTreatedAsLax
	case SameSiteCookieExclusionReasonExcludeSameSiteNoneInsecure:
		*t = SameSiteCookieExclusionReasonExcludeSameSiteNoneInsecure

	default:
		in.AddError(errors.New("unknown SameSiteCookieExclusionReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SameSiteCookieExclusionReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SameSiteCookieWarningReason [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieWarningReason
type SameSiteCookieWarningReason string

// String returns the SameSiteCookieWarningReason as string value.
func (t SameSiteCookieWarningReason) String() string {
	return string(t)
}

// SameSiteCookieWarningReason values.
const (
	SameSiteCookieWarningReasonWarnSameSiteUnspecifiedCrossSiteContext        SameSiteCookieWarningReason = "WarnSameSiteUnspecifiedCrossSiteContext"
	SameSiteCookieWarningReasonWarnSameSiteNoneInsecure                       SameSiteCookieWarningReason = "WarnSameSiteNoneInsecure"
	SameSiteCookieWarningReasonWarnSameSiteUnspecifiedLaxAllowUnsafe          SameSiteCookieWarningReason = "WarnSameSiteUnspecifiedLaxAllowUnsafe"
	SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLMethodUnsafe   SameSiteCookieWarningReason = "WarnSameSiteCrossSchemeSecureUrlMethodUnsafe"
	SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLLax            SameSiteCookieWarningReason = "WarnSameSiteCrossSchemeSecureUrlLax"
	SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLStrict         SameSiteCookieWarningReason = "WarnSameSiteCrossSchemeSecureUrlStrict"
	SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLMethodUnsafe SameSiteCookieWarningReason = "WarnSameSiteCrossSchemeInsecureUrlMethodUnsafe"
	SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLLax          SameSiteCookieWarningReason = "WarnSameSiteCrossSchemeInsecureUrlLax"
	SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLStrict       SameSiteCookieWarningReason = "WarnSameSiteCrossSchemeInsecureUrlStrict"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SameSiteCookieWarningReason) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SameSiteCookieWarningReason) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SameSiteCookieWarningReason) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SameSiteCookieWarningReason(in.String()) {
	case SameSiteCookieWarningReasonWarnSameSiteUnspecifiedCrossSiteContext:
		*t = SameSiteCookieWarningReasonWarnSameSiteUnspecifiedCrossSiteContext
	case SameSiteCookieWarningReasonWarnSameSiteNoneInsecure:
		*t = SameSiteCookieWarningReasonWarnSameSiteNoneInsecure
	case SameSiteCookieWarningReasonWarnSameSiteUnspecifiedLaxAllowUnsafe:
		*t = SameSiteCookieWarningReasonWarnSameSiteUnspecifiedLaxAllowUnsafe
	case SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLMethodUnsafe:
		*t = SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLMethodUnsafe
	case SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLLax:
		*t = SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLLax
	case SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLStrict:
		*t = SameSiteCookieWarningReasonWarnSameSiteCrossSchemeSecureURLStrict
	case SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLMethodUnsafe:
		*t = SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLMethodUnsafe
	case SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLLax:
		*t = SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLLax
	case SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLStrict:
		*t = SameSiteCookieWarningReasonWarnSameSiteCrossSchemeInsecureURLStrict

	default:
		in.AddError(errors.New("unknown SameSiteCookieWarningReason value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SameSiteCookieWarningReason) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SameSiteCookieOperation [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieOperation
type SameSiteCookieOperation string

// String returns the SameSiteCookieOperation as string value.
func (t SameSiteCookieOperation) String() string {
	return string(t)
}

// SameSiteCookieOperation values.
const (
	SameSiteCookieOperationSetCookie  SameSiteCookieOperation = "SetCookie"
	SameSiteCookieOperationReadCookie SameSiteCookieOperation = "ReadCookie"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t SameSiteCookieOperation) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t SameSiteCookieOperation) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *SameSiteCookieOperation) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch SameSiteCookieOperation(in.String()) {
	case SameSiteCookieOperationSetCookie:
		*t = SameSiteCookieOperationSetCookie
	case SameSiteCookieOperationReadCookie:
		*t = SameSiteCookieOperationReadCookie

	default:
		in.AddError(errors.New("unknown SameSiteCookieOperation value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *SameSiteCookieOperation) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// SameSiteCookieIssueDetails this information is currently necessary, as the
// front-end has a difficult time finding a specific cookie. With this, we can
// convey specific error information without the cookie.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-SameSiteCookieIssueDetails
type SameSiteCookieIssueDetails struct {
	Cookie                 *AffectedCookie                 `json:"cookie"`
	CookieWarningReasons   []SameSiteCookieWarningReason   `json:"cookieWarningReasons"`
	CookieExclusionReasons []SameSiteCookieExclusionReason `json:"cookieExclusionReasons"`
	Operation              SameSiteCookieOperation         `json:"operation"` // Optionally identifies the site-for-cookies and the cookie url, which may be used by the front-end as additional context.
	SiteForCookies         string                          `json:"siteForCookies,omitempty"`
	CookieURL              string                          `json:"cookieUrl,omitempty"`
	Request                *AffectedRequest                `json:"request,omitempty"`
}

// MixedContentResolutionStatus [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-MixedContentResolutionStatus
type MixedContentResolutionStatus string

// String returns the MixedContentResolutionStatus as string value.
func (t MixedContentResolutionStatus) String() string {
	return string(t)
}

// MixedContentResolutionStatus values.
const (
	MixedContentResolutionStatusMixedContentBlocked               MixedContentResolutionStatus = "MixedContentBlocked"
	MixedContentResolutionStatusMixedContentAutomaticallyUpgraded MixedContentResolutionStatus = "MixedContentAutomaticallyUpgraded"
	MixedContentResolutionStatusMixedContentWarning               MixedContentResolutionStatus = "MixedContentWarning"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MixedContentResolutionStatus) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t MixedContentResolutionStatus) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MixedContentResolutionStatus) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch MixedContentResolutionStatus(in.String()) {
	case MixedContentResolutionStatusMixedContentBlocked:
		*t = MixedContentResolutionStatusMixedContentBlocked
	case MixedContentResolutionStatusMixedContentAutomaticallyUpgraded:
		*t = MixedContentResolutionStatusMixedContentAutomaticallyUpgraded
	case MixedContentResolutionStatusMixedContentWarning:
		*t = MixedContentResolutionStatusMixedContentWarning

	default:
		in.AddError(errors.New("unknown MixedContentResolutionStatus value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MixedContentResolutionStatus) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// MixedContentResourceType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-MixedContentResourceType
type MixedContentResourceType string

// String returns the MixedContentResourceType as string value.
func (t MixedContentResourceType) String() string {
	return string(t)
}

// MixedContentResourceType values.
const (
	MixedContentResourceTypeAudio          MixedContentResourceType = "Audio"
	MixedContentResourceTypeBeacon         MixedContentResourceType = "Beacon"
	MixedContentResourceTypeCSPReport      MixedContentResourceType = "CSPReport"
	MixedContentResourceTypeDownload       MixedContentResourceType = "Download"
	MixedContentResourceTypeEventSource    MixedContentResourceType = "EventSource"
	MixedContentResourceTypeFavicon        MixedContentResourceType = "Favicon"
	MixedContentResourceTypeFont           MixedContentResourceType = "Font"
	MixedContentResourceTypeForm           MixedContentResourceType = "Form"
	MixedContentResourceTypeFrame          MixedContentResourceType = "Frame"
	MixedContentResourceTypeImage          MixedContentResourceType = "Image"
	MixedContentResourceTypeImport         MixedContentResourceType = "Import"
	MixedContentResourceTypeManifest       MixedContentResourceType = "Manifest"
	MixedContentResourceTypePing           MixedContentResourceType = "Ping"
	MixedContentResourceTypePluginData     MixedContentResourceType = "PluginData"
	MixedContentResourceTypePluginResource MixedContentResourceType = "PluginResource"
	MixedContentResourceTypePrefetch       MixedContentResourceType = "Prefetch"
	MixedContentResourceTypeResource       MixedContentResourceType = "Resource"
	MixedContentResourceTypeScript         MixedContentResourceType = "Script"
	MixedContentResourceTypeServiceWorker  MixedContentResourceType = "ServiceWorker"
	MixedContentResourceTypeSharedWorker   MixedContentResourceType = "SharedWorker"
	MixedContentResourceTypeStylesheet     MixedContentResourceType = "Stylesheet"
	MixedContentResourceTypeTrack          MixedContentResourceType = "Track"
	MixedContentResourceTypeVideo          MixedContentResourceType = "Video"
	MixedContentResourceTypeWorker         MixedContentResourceType = "Worker"
	MixedContentResourceTypeXMLHTTPRequest MixedContentResourceType = "XMLHttpRequest"
	MixedContentResourceTypeXSLT           MixedContentResourceType = "XSLT"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t MixedContentResourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t MixedContentResourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *MixedContentResourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch MixedContentResourceType(in.String()) {
	case MixedContentResourceTypeAudio:
		*t = MixedContentResourceTypeAudio
	case MixedContentResourceTypeBeacon:
		*t = MixedContentResourceTypeBeacon
	case MixedContentResourceTypeCSPReport:
		*t = MixedContentResourceTypeCSPReport
	case MixedContentResourceTypeDownload:
		*t = MixedContentResourceTypeDownload
	case MixedContentResourceTypeEventSource:
		*t = MixedContentResourceTypeEventSource
	case MixedContentResourceTypeFavicon:
		*t = MixedContentResourceTypeFavicon
	case MixedContentResourceTypeFont:
		*t = MixedContentResourceTypeFont
	case MixedContentResourceTypeForm:
		*t = MixedContentResourceTypeForm
	case MixedContentResourceTypeFrame:
		*t = MixedContentResourceTypeFrame
	case MixedContentResourceTypeImage:
		*t = MixedContentResourceTypeImage
	case MixedContentResourceTypeImport:
		*t = MixedContentResourceTypeImport
	case MixedContentResourceTypeManifest:
		*t = MixedContentResourceTypeManifest
	case MixedContentResourceTypePing:
		*t = MixedContentResourceTypePing
	case MixedContentResourceTypePluginData:
		*t = MixedContentResourceTypePluginData
	case MixedContentResourceTypePluginResource:
		*t = MixedContentResourceTypePluginResource
	case MixedContentResourceTypePrefetch:
		*t = MixedContentResourceTypePrefetch
	case MixedContentResourceTypeResource:
		*t = MixedContentResourceTypeResource
	case MixedContentResourceTypeScript:
		*t = MixedContentResourceTypeScript
	case MixedContentResourceTypeServiceWorker:
		*t = MixedContentResourceTypeServiceWorker
	case MixedContentResourceTypeSharedWorker:
		*t = MixedContentResourceTypeSharedWorker
	case MixedContentResourceTypeStylesheet:
		*t = MixedContentResourceTypeStylesheet
	case MixedContentResourceTypeTrack:
		*t = MixedContentResourceTypeTrack
	case MixedContentResourceTypeVideo:
		*t = MixedContentResourceTypeVideo
	case MixedContentResourceTypeWorker:
		*t = MixedContentResourceTypeWorker
	case MixedContentResourceTypeXMLHTTPRequest:
		*t = MixedContentResourceTypeXMLHTTPRequest
	case MixedContentResourceTypeXSLT:
		*t = MixedContentResourceTypeXSLT

	default:
		in.AddError(errors.New("unknown MixedContentResourceType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *MixedContentResourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// MixedContentIssueDetails [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-MixedContentIssueDetails
type MixedContentIssueDetails struct {
	ResourceType     MixedContentResourceType     `json:"resourceType,omitempty"` // The type of resource causing the mixed content issue (css, js, iframe, form,...). Marked as optional because it is mapped to from blink::mojom::RequestContextType, which will be replaced by network::mojom::RequestDestination
	ResolutionStatus MixedContentResolutionStatus `json:"resolutionStatus"`       // The way the mixed content issue is being resolved.
	InsecureURL      string                       `json:"insecureURL"`            // The unsafe http url causing the mixed content issue.
	MainResourceURL  string                       `json:"mainResourceURL"`        // The url responsible for the call to an unsafe url.
	Request          *AffectedRequest             `json:"request,omitempty"`      // The mixed content request. Does not always exist (e.g. for unsafe form submission urls).
	Frame            *AffectedFrame               `json:"frame,omitempty"`        // Optional because not every mixed content issue is necessarily linked to a frame.
}

// InspectorIssueCode a unique identifier for the type of issue. Each type
// may use one of the optional fields in InspectorIssueDetails to convey more
// specific information about the kind of issue.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-InspectorIssueCode
type InspectorIssueCode string

// String returns the InspectorIssueCode as string value.
func (t InspectorIssueCode) String() string {
	return string(t)
}

// InspectorIssueCode values.
const (
	InspectorIssueCodeSameSiteCookieIssue InspectorIssueCode = "SameSiteCookieIssue"
	InspectorIssueCodeMixedContentIssue   InspectorIssueCode = "MixedContentIssue"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t InspectorIssueCode) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t InspectorIssueCode) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *InspectorIssueCode) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch InspectorIssueCode(in.String()) {
	case InspectorIssueCodeSameSiteCookieIssue:
		*t = InspectorIssueCodeSameSiteCookieIssue
	case InspectorIssueCodeMixedContentIssue:
		*t = InspectorIssueCodeMixedContentIssue

	default:
		in.AddError(errors.New("unknown InspectorIssueCode value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *InspectorIssueCode) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// InspectorIssueDetails this struct holds a list of optional fields with
// additional information specific to the kind of issue. When adding a new issue
// code, please also add a new optional field to this type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-InspectorIssueDetails
type InspectorIssueDetails struct {
	SameSiteCookieIssueDetails *SameSiteCookieIssueDetails `json:"sameSiteCookieIssueDetails,omitempty"`
	MixedContentIssueDetails   *MixedContentIssueDetails   `json:"mixedContentIssueDetails,omitempty"`
}

// InspectorIssue an inspector issue reported from the back-end.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#type-InspectorIssue
type InspectorIssue struct {
	Code    InspectorIssueCode     `json:"code"`
	Details *InspectorIssueDetails `json:"details"`
}

// GetEncodedResponseEncoding the encoding to use.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Audits#method-getEncodedResponse
type GetEncodedResponseEncoding string

// String returns the GetEncodedResponseEncoding as string value.
func (t GetEncodedResponseEncoding) String() string {
	return string(t)
}

// GetEncodedResponseEncoding values.
const (
	GetEncodedResponseEncodingWebp GetEncodedResponseEncoding = "webp"
	GetEncodedResponseEncodingJpeg GetEncodedResponseEncoding = "jpeg"
	GetEncodedResponseEncodingPng  GetEncodedResponseEncoding = "png"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t GetEncodedResponseEncoding) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t GetEncodedResponseEncoding) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *GetEncodedResponseEncoding) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch GetEncodedResponseEncoding(in.String()) {
	case GetEncodedResponseEncodingWebp:
		*t = GetEncodedResponseEncodingWebp
	case GetEncodedResponseEncodingJpeg:
		*t = GetEncodedResponseEncodingJpeg
	case GetEncodedResponseEncodingPng:
		*t = GetEncodedResponseEncodingPng

	default:
		in.AddError(errors.New("unknown GetEncodedResponseEncoding value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *GetEncodedResponseEncoding) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
