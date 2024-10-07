package api

import (
    "context"
    "fmt"
    "goblossom/internal/blob"
    "goblossom/internal/api"
    "goblossom/pkg/utils"
)

// HandlerImpl is the implementation of the ogen-go generated Handler interface
type HandlerImpl struct {
    BlobRepo blob.BlobRepository // Inject BlobRepository to handle blob operations
}

// ListPubkeyGet handles GET /list/{pubkey}
func (h *HandlerImpl) ListPubkeyGet(ctx context.Context, params api.ListPubkeyGetParams) (api.ListPubkeyGetRes, error) {
    blobs, err := h.BlobRepo.ListBlobsByPubKey(params.Pubkey, 0, 0)
    if err != nil {
        return nil, fmt.Errorf("failed to list blobs: %v", err)
    }
    return blobs, nil
}

// MirrorPut handles PUT /mirror
func (h *HandlerImpl) MirrorPut(ctx context.Context, req *api.MirrorPutReq) (api.MirrorPutRes, error) {
    err := h.BlobRepo.MirrorBlob(req.Url)
    if err != nil {
        return nil, fmt.Errorf("failed to mirror blob: %v", err)
    }
    return api.MirrorPutRes{}, nil
}

// SHA256Delete handles DELETE /{sha256}
func (h *HandlerImpl) SHA256Delete(ctx context.Context, params api.SHA256DeleteParams) (api.SHA256DeleteRes, error) {
    err := h.BlobRepo.DeleteBlob(params.SHA256)
    if err != nil {
        return nil, fmt.Errorf("failed to delete blob: %v", err)
    }
    return api.SHA256DeleteRes{}, nil
}

// SHA256Get handles GET /{sha256}
func (h *HandlerImpl) SHA256Get(ctx context.Context, params api.SHA256GetParams) (api.SHA256GetRes, error) {
    blobData, descriptor, err := h.BlobRepo.GetBlob(params.SHA256)
    if err != nil {
        return nil, fmt.Errorf("failed to get blob: %v", err)
    }
    return api.SHA256GetRes{
        Data:        blobData,
        ContentType: descriptor.Type,
    }, nil
}

// SHA256Head handles HEAD /{sha256}
func (h *HandlerImpl) SHA256Head(ctx context.Context, params api.SHA256HeadParams) (api.SHA256HeadRes, error) {
    exists, err := h.BlobRepo.HasBlob(params.Sha256)
    if err != nil || !exists {
        return nil, fmt.Errorf("blob not found: %v", err)
    }
    return api.SHA256HeadRes{Exists: exists}, nil
}

// UploadHead handles HEAD /upload
func (h *HandlerImpl) UploadHead(ctx context.Context, params api.UploadHeadParams) (api.UploadHeadRes, error) {
    valid := h.BlobRepo.IsAllowedFileType(params.XContentType)
    if !valid {
        return nil, fmt.Errorf("unsupported file type: %v", params.XContentType)
    }
    return UploadHeadRes{}, nil
}

// UploadPut handles PUT /upload
func (h *HandlerImpl) UploadPut(ctx context.Context, req api.UploadPutReq) (api.UploadPutRes, error) {
    // Read file data from the request body
    hash, err := utils.HashReader(req.File)
    if err != nil {
        return nil, fmt.Errorf("failed to hash file: %v", err)
    }

    // Save the blob using BlobRepository
    descriptor, err := h.BlobRepo.SaveBlob(hash, req.File, req.Size, req.XContentType)
    if err != nil {
        return nil, fmt.Errorf("failed to save blob: %v", err)
    }

    return api.UploadPutRes{
        Sha256: hash,
        Url:    descriptor.URL,
        Size:   descriptor.Size,
        Type:   descriptor.Type,
    }, nil
}
