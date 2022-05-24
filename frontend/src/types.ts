export interface Post {
    Id: number
    Content: string
    Poster?: string
    IsLocked: boolean
    CreatedAt: Date
}

export interface Posts {
    CurrentPage: number
    NextPage?: number
    PrevPage?: number
    Results: Post[]
    Total: number
}